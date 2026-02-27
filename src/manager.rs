use crate::models::{TodoItem, TodoStatus};

pub struct TodoManager {
    pub todos: Vec<TodoItem>,
    pub next_id: usize,
    pub free_ids: Vec<usize>,
}

impl TodoManager {
    pub fn new() -> Self {
        TodoManager {
            todos: Vec::new(),
            next_id: 1,
            free_ids: Vec::new(),
        }
    }

    pub fn add_todo(&mut self, title: String) {
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

    pub fn delete_todo(&mut self, id: usize) -> bool {
        let len_before = self.todos.len();
        self.todos.retain(|t| t.id != id);
        if self.todos.len() < len_before {
            self.free_ids.push(id);
            return true;
        }
        false
    }

    pub fn complete_todo(&mut self, id: usize) -> bool {
        for todo in &mut self.todos {
            if todo.id == id {
                todo.status = TodoStatus::Completed;
                return true;
            }
        }
        false
    }

    pub fn list_todos(&self) {
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

    pub fn save_to_file(&self, path: &str) -> std::io::Result<()> {
        if let Some(parent) = std::path::Path::new(path).parent() {
            std::fs::create_dir_all(parent)?;
        }
        let s = serde_json::to_string_pretty(&self.todos).unwrap();
        std::fs::write(path, s)
    }

    pub fn load_from_file(path: &str) -> Self {
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
