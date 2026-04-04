# Go REST API — Daily Planner

Phase 2A of the [DailyPlanner monorepo](../README.md). A lightweight REST API built with Go's standard library.

---

## Status

🚧 **In Progress**

---

## Quick Start

**Prerequisites:** Go 1.21+

```bash
cd go-api

# Run the server
go run ./cmd/main.go

# Run tests (once added)
go test ./...
```

Server starts on `http://localhost:8080`.

---

## Project Structure

```
go-api/
├── cmd/
│   └── main.go           # Entry point, route registration
├── internal/
│   ├── models/
│   │   └── task.go       # Task struct
│   └── handlers/
│       └── tasks.go      # HTTP handler functions
└── go.mod
```

---

## Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/health` | Health check |
| `GET` | `/tasks` | List all tasks |
| `POST` | `/tasks` | Create a task |
| `DELETE` | `/tasks/{id}` | Delete a task |
| `PUT` | `/tasks/{id}/complete` | Mark a task complete |

---

## Learning Goals

- Go module system and package layout
- `net/http` standard library (no framework)
- JSON marshaling/unmarshaling (`encoding/json`)
- HTTP method routing (`r.Method`)
- Go maps and in-memory state
- SQLite persistence (upcoming)
- Goroutines and concurrency patterns (upcoming)

---

## Key Go Concepts Covered

| Concept | Rust Equivalent |
|---------|----------------|
| `package` / `import` | `mod` / `use` |
| Capitalized name = exported | `pub` keyword |
| `struct` with field tags | `struct` with serde attributes |
| `[]T` slice | `Vec<T>` |
| `map[K]V` | `HashMap<K, V>` |
| `encoding/json` | `serde_json` |

---

## CI

CI workflow: [../.github/workflows/go-api-ci.yml](../.github/workflows/go-api-ci.yml) *(coming soon)*
