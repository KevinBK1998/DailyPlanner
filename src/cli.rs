use crate::manager::TodoManager;
use crate::models::Command;

pub fn parse_id(arg: Option<&str>) -> Result<usize, &'static str> {
    match arg {
        Some(s) => s.parse::<usize>().map_err(|_| "id must be a number"),
        None => Err("id is required"),
    }
}

pub fn parse_command(input: &str) -> Result<Command, &'static str> {
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

pub fn save_or_print(mgr: &TodoManager, path: &str) {
    if let Err(e) = mgr.save_to_file(path) {
        eprintln!("save error: {}", e);
    }
}

pub fn execute_command(cmd: Command, mgr: &mut TodoManager, path: &str) -> bool {
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

#[cfg(test)]
mod tests {
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
}
