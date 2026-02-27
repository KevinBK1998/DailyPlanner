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

#[derive(PartialEq, Debug)]
enum Command {
    Add(String),
    Delete(usize),
    Complete(usize),
    List,
    Exit,
    Help,
}

fn parse_command(input: &str) -> Result<Command, &'static str> {
    let parts: Vec<_> = input.trim().splitn(2, ' ').collect();
    match parts[0] {
        "add" => {
            let title = parts.get(1).copied().ok_or("title is required")?;
            Ok(Command::Add(title.to_string()))
        }
        "delete" => Ok(Command::Delete(parse_id(parts.get(1).copied())?)),
        "complete" => Ok(Command::Complete(parse_id(parts.get(1).copied())?)),
        "list" => Ok(Command::List),
        "exit" => Ok(Command::Exit),
        "help" => Ok(Command::Help),
        _ => Err("unknown command"),
    }
}

fn execute_command(cmd: Command, mgr: &mut TodoManager, path: &str) -> bool {
    match cmd {
        Command::Add(title) => {
            mgr.add_todo(title);
            save_or_print(mgr, path);
            true
        }
        Command::Delete(id) => {
            if !mgr.delete_todo(id) {
                println!("no such id");
            }
            save_or_print(mgr, path);
            true
        }
        Command::Complete(id) => {
            if !mgr.complete_todo(id) {
                println!("no such id");
            }
            save_or_print(mgr, path);
            true
        }
        Command::List => {
            mgr.list_todos();
            true
        }
        Command::Help => {
            println!("Commands: add <title>, complete <id>, delete <id>, list, help, exit");
            true
        }
        Command::Exit => false,
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
        match parse_command(&input) {
            Ok(cmd) => {
                if !execute_command(cmd, &mut mgr, path) {
                    break;
                }
            }
            Err(msg) => println!("{}", msg),
        }
    }
}

#[cfg(test)]
mod tests {
    fn temp_test_file(name: &str) -> std::path::PathBuf {
        let mut p = std::env::temp_dir();
        let stamp = std::time::SystemTime::now()
            .duration_since(std::time::UNIX_EPOCH)
            .unwrap()
            .as_nanos();
        p.push(format!("daily_planner_{}_{}.json", name, stamp));
        p
    }
    use super::*;

    #[test]
    fn parse_id_returns_ok_for_number() {
        let result = parse_id(Some("42"));
        assert_eq!(result, Ok(42));
    }

    #[test]
    fn parse_id_errors_for_missing_arg() {
        let result = parse_id(None);
        assert_eq!(result, Err("id is required"));
    }

    #[test]
    fn parse_id_errors_for_non_numeric_arg() {
        let result = parse_id(Some("abc"));
        assert_eq!(result, Err("id must be a number"));
    }

    #[test]
    fn parse_command_add_returns_add_variant() {
        let cmd = parse_command("add buy milk");
        match cmd {
            Ok(Command::Add(title)) => assert_eq!(title, "buy milk"),
            _ => panic!("expected Command::Add"),
        }
    }

    #[test]
    fn parse_command_delete_returns_delete_variant() {
        let cmd = parse_command("delete 7");
        match cmd {
            Ok(Command::Delete(id)) => assert_eq!(id, 7),
            _ => panic!("expected Command::Delete"),
        }
    }

    #[test]
    fn parse_command_unknown_returns_error() {
        let cmd = parse_command("wat");
        assert_eq!(cmd, Err("unknown command"));
    }

    #[test]
    fn parse_command_add_missing_title_returns_error() {
        let cmd = parse_command("add");
        assert_eq!(cmd, Err("title is required"));
    }

    #[test]
    fn parse_command_delete_non_numeric_id_returns_error() {
        let cmd = parse_command("delete abc");
        assert_eq!(cmd, Err("id must be a number"));
    }

    #[test]
    fn add_todo_increments_id_and_stores_item() {
        let mut mgr = TodoManager::new();
        mgr.add_todo("first".to_string());
        mgr.add_todo("second".to_string());

        assert_eq!(mgr.todos.len(), 2);
        assert_eq!(mgr.todos[0].id, 1);
        assert_eq!(mgr.todos[1].id, 2);
        assert_eq!(mgr.next_id, 3);
    }

    #[test]
    fn delete_todo_reuses_deleted_id() {
        let mut mgr = TodoManager::new();
        mgr.add_todo("a".to_string()); // id 1
        mgr.add_todo("b".to_string()); // id 2

        assert!(mgr.delete_todo(1));
        mgr.add_todo("c".to_string()); // should reuse id 1

        let ids: Vec<usize> = mgr.todos.iter().map(|t| t.id).collect();
        assert!(ids.contains(&1));
        assert!(ids.contains(&2));
    }

    #[test]
    fn complete_todo_marks_status_completed() {
        let mut mgr = TodoManager::new();
        mgr.add_todo("x".to_string());

        assert!(mgr.complete_todo(1));
        match mgr.todos[0].status {
            TodoStatus::Completed => {}
            _ => panic!("expected Completed"),
        }
    }

    #[test]
    fn save_and_load_round_trip_preserves_todos() {
        let path = temp_test_file("round_trip");
        let path_str = path.to_str().unwrap();

        let mut mgr = TodoManager::new();
        mgr.add_todo("one".to_string());
        mgr.add_todo("two".to_string());
        mgr.complete_todo(2);

        mgr.save_to_file(path_str).unwrap();

        let loaded = TodoManager::load_from_file(path_str);
        assert_eq!(loaded.todos.len(), 2);
        assert_eq!(
            loaded.todos.iter().map(|t| t.id).collect::<Vec<_>>(),
            vec![1, 2]
        );
        let _ = std::fs::remove_file(&path);
    }

    #[test]
    fn load_reconstructs_free_ids_from_gaps() {
        let path = temp_test_file("free_id_trip");
        let path_str = path.to_str().unwrap();

        std::fs::write(path_str, r#"[{"id":2,"title":"only","status":"Pending"}]"#).unwrap();

        let mut loaded = TodoManager::load_from_file(path_str);
        loaded.add_todo("reuse".to_string());

        let ids: Vec<usize> = loaded.todos.iter().map(|t| t.id).collect();
        assert!(ids.contains(&1));
        assert!(ids.contains(&2));
        let _ = std::fs::remove_file(&path);
    }
}
