enum TodoStatus {
    Pending,
    Completed,
}
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
}
fn main() {
    println!("Welcome to Daily Planner!");

    // create a manager and exercise its methods
    let mut manager = TodoManager::new();
    manager.add_todo("Learn Rust".to_string());
    manager.add_todo("Write a todo app".to_string());

    // list current todos (demonstrates borrowing)
    manager.list_todos();
}
