use std::io::{self, Write};

mod cli;
mod manager;
mod models;
use cli::{execute_command, parse_command};
use manager::TodoManager;

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
