use std::io::{self, Write};

use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize)]
enum TodoStatus {
    Pending,
    Completed,
}

#[derive(Serialize, Deserialize)]
struct TodoItem {
    id: usize,
    title: String,
    status: TodoStatus,
}
struct TodoManager {
    todos: Vec<TodoItem>,
    next_id: usize,
    free_ids: Vec<usize>,
}

impl TodoManager {
    fn new() -> Self {
        TodoManager {
            todos: Vec::new(),
            next_id: 1,
            free_ids: Vec::new(),
        }
    }
    fn add_todo(&mut self, title: String) {
        let id = self.free_ids.pop().unwrap_or_else(|| {
            let id = self.next_id;
            self.next_id += 1;
            id
        });
        let todo = TodoItem {
            id,
            title,
            status: TodoStatus::Pending,
        };
        self.todos.push(todo);
    }
    fn delete_todo(&mut self, id: usize) -> bool {
        let len_before = self.todos.len();
        self.todos.retain(|t| t.id != id);
        if self.todos.len() < len_before {
            self.free_ids.push(id);
            return true;
        }
        false
    }
    fn complete_todo(&mut self, id: usize) -> bool {
        for todo in &mut self.todos {
            if todo.id == id {
                todo.status = TodoStatus::Completed;
                return true;
            }
        }
        false
    }
    fn list_todos(&self) {
        if self.todos.is_empty() {
            println!("No todos yet!");
            return;
        }
        for todo in &self.todos {
            let status = match todo.status {
                TodoStatus::Completed => "Completed",
                TodoStatus::Pending => "Pending",
            };
            println!("ID: {}, Title: {}, Status: {}", todo.id, todo.title, status);
        }
    }
    fn save_to_file(&self, path: &str) -> std::io::Result<()> {
        if let Some(parent) = std::path::Path::new(path).parent() {
            std::fs::create_dir_all(parent)?;
        }
        let s = serde_json::to_string_pretty(&self.todos).unwrap();
        std::fs::write(path, s)
    }
    fn load_from_file(path: &str) -> Self {
        match std::fs::read_to_string(path) {
            Ok(data) => {
                let todos: Vec<TodoItem> = serde_json::from_str(&data).unwrap_or_default();
                let max_id = todos.iter().map(|t| t.id).max().unwrap_or(0);
                let mut used = vec![false; max_id + 1];
                for todo in &todos {
                    if todo.id <= max_id {
                        used[todo.id] = true;
                    }
                }
                let mut free_ids = Vec::new();
                for id in (1..=max_id).rev() {
                    if !used[id] {
                        free_ids.push(id);
                    }
                }
                let next_id = max_id + 1;
                TodoManager {
                    todos,
                    next_id,
                    free_ids,
                }
            }
            Err(_) => TodoManager::new(),
        }
    }
}

fn parse_id(arg: Option<&str>) -> Result<usize, &'static str> {
    match arg {
        Some(s) => s.parse::<usize>().map_err(|_| "id must be a number"),
        None => Err("id is required"),
    }
}

fn save_or_print(mgr: &TodoManager, path: &str) {
    if let Err(e) = mgr.save_to_file(path) {
        eprintln!("save error: {}", e);
    }
}

fn main() {
    println!("Welcome to Daily Planner!");
    let path = "data/todos.json";
    let mut mgr = TodoManager::load_from_file(path);
    loop {
        print!("cmd> ");
        io::stdout().flush().unwrap();
        let mut input = String::new();
        io::stdin().read_line(&mut input).unwrap();
        let parts: Vec<_> = input.trim().splitn(2, ' ').collect();
        match parts[0] {
            "add" if parts.len() == 2 => {
                mgr.add_todo(parts[1].to_string());
                save_or_print(&mgr, path);
            }
            "delete" => match parse_id(parts.get(1).copied()) {
                Ok(id) => {
                    if !mgr.delete_todo(id) {
                        println!("no such id");
                    }
                    save_or_print(&mgr, path);
                }
                Err(msg) => println!("{}", msg),
            },
            "complete" => match parse_id(parts.get(1).copied()) {
                Ok(id) => {
                    if !mgr.complete_todo(id) {
                        println!("no such id");
                    }
                    save_or_print(&mgr, path);
                }
                Err(msg) => println!("{}", msg),
            },
            "list" => {
                mgr.list_todos();
            }
            "exit" => break,
            _ => println!("Commands: add <title>, complete <id>, delete <id>, exit, list"),
        }
    }
}
