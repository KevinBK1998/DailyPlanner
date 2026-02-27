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
}

impl TodoManager {
    fn new() -> Self {
        TodoManager {
            todos: Vec::new(),
            next_id: 1,
        }
    }
    fn add_todo(&mut self, title: String) {
        let todo = TodoItem {
            id: self.next_id,
            title,
            status: TodoStatus::Pending,
        };
        self.todos.push(todo);
        self.next_id += 1;
    }
    fn list_todos(&self) {
        if self.todos.is_empty() {
            println!("No todos yet!");
            return;
        }
        for todo in &self.todos {
            println!("ID: {}, Title: {}", todo.id, todo.title);
        }
    }
    fn save_to_file(&self, path: &str) -> std::io::Result<()> {
        let s = serde_json::to_string_pretty(&self.todos).unwrap();
        std::fs::write(path, s)
    }
    fn load_from_file(path: &str) -> Self {
        match std::fs::read_to_string(path) {
            Ok(data) => {
                let todos: Vec<TodoItem> = serde_json::from_str(&data).unwrap_or_default();
                let next_id = todos.iter().map(|t| t.id).max().unwrap_or(0) + 1;
                TodoManager { todos, next_id }
            }
            Err(_) => TodoManager::new(),
        }
    }
}

fn main() {
    println!("Welcome to Daily Planner!");
    let mut mgr = TodoManager::load_from_file("todos.json");
    loop {
        print!("cmd> ");
        io::stdout().flush().unwrap();
        let mut input = String::new();
        io::stdin().read_line(&mut input).unwrap();
        let parts: Vec<_> = input.trim().splitn(2, ' ').collect();
        match parts[0] {
            "init" => {
                mgr.add_todo("Learn Rust".to_string());
                mgr.add_todo("Write a todo app".to_string());
            }
            "list" => {
                mgr.list_todos();
            }
            "save" => {
                mgr.save_to_file("todos.json");
            }
            "exit" => break,
            _ => println!("Commands: init, list, save"),
        }
    }
}
