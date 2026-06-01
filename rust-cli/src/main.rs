use std::io::{self, Write};

mod cli;
mod manager;
mod models;
use cli::{execute_command, parse_command};
use manager::TodoManager;

fn flush_prompt<W: Write>(writer: &mut W) -> io::Result<()> {
    write!(writer, "cmd> ")?;
    writer.flush()
}

fn read_input_line() -> io::Result<String> {
    let mut input = String::new();
    io::stdin().read_line(&mut input)?;
    Ok(input)
}

fn main() {
    println!("Welcome to Daily Planner!");
    let path = "data/todos.json";
    let mut mgr = TodoManager::load_from_file(path);

    loop {
        let mut stdout = io::stdout();
        if let Err(err) = flush_prompt(&mut stdout) {
            eprintln!("prompt error: {}", err);
            break;
        }

        let input = match read_input_line() {
            Ok(value) => value,
            Err(err) => {
                eprintln!("input error: {}", err);
                break;
            }
        };

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
    use super::*;

    struct FailingWriter;

    impl Write for FailingWriter {
        fn write(&mut self, _buf: &[u8]) -> io::Result<usize> {
            Err(io::Error::other("write failed"))
        }

        fn flush(&mut self) -> io::Result<()> {
            Err(io::Error::other("flush failed"))
        }
    }

    #[test]
    fn flush_prompt_writes_prompt() {
        let mut buf = Vec::new();
        let result = flush_prompt(&mut buf);
        assert!(result.is_ok());
        assert_eq!(buf, b"cmd> ");
    }

    #[test]
    fn flush_prompt_returns_error_when_writer_fails() {
        let mut writer = FailingWriter;
        let result = flush_prompt(&mut writer);
        assert!(result.is_err());
    }
}
