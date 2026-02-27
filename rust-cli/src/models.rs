use serde::{Deserialize, Serialize};

#[derive(Debug, PartialEq, Serialize, Deserialize)]
pub enum TodoStatus {
    Pending,
    Completed,
}

#[derive(Debug, PartialEq, Serialize, Deserialize)]
pub struct TodoItem {
    pub id: usize,
    pub title: String,
    pub status: TodoStatus,
}

#[derive(PartialEq, Debug)]
pub enum Command {
    Add(String),
    Delete(usize),
    Complete(usize),
    List,
    Exit,
    Help,
}
