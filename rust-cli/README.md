# 🦀 Daily Planner Rust CLI

Rust implementation of the Phase 1 todo manager.

## Status

- Phase 1 (Rust CLI): ✅ Completed
- Current capabilities: add, list, complete, delete, JSON persistence
- Test suite: parser tests + manager behavior tests + persistence tests

## Quick Start

### Build
```bash
cargo build
```

### Run
```bash
cargo run
```

### Test
```bash
cargo test
```

## Project Structure

```
rust-cli/
├── Cargo.toml
├── Cargo.lock
├── data/
│   └── todos.json
└── src/
    ├── main.rs      # runtime wiring + input loop
    ├── cli.rs       # parse + execute commands
    ├── manager.rs   # TodoManager domain logic
    └── models.rs    # TodoStatus, TodoItem, Command
```

## Commands

- `add <title>`
- `list`
- `complete <id>`
- `delete <id>`
- `help`
- `exit`

## Release Workflow

Automated release is configured in:

- `../.github/workflows/rust-cli-release.yml`

Release trigger tag format:

- `rust-cli-v*` (example: `rust-cli-v0.1.0`)

### Create a release

```bash
git tag rust-cli-v0.1.0
git push origin rust-cli-v0.1.0
```

The workflow will build and attach binaries for Linux and Windows to a GitHub Release.

## CI Workflow

Rust CI is configured in:

- `../.github/workflows/rust-cli-ci.yml`

It runs on pushes/PRs touching `rust-cli/**` and performs:

- `cargo fmt --check`
- `cargo clippy -D warnings`
- `cargo test`