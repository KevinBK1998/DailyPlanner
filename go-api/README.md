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

# Build + compile-check all packages
go test ./...
```

Server starts on `http://localhost:8080`.

---

## Current Status

Implemented and verified:

- `GET /health`
- `GET /tasks`
- `POST /tasks`
- `DELETE /tasks/{id}`
- `PUT /tasks/{id}/complete`
- In-memory task store (`internal/store/task_store.go`)
- Bruno request collection (`bruno/requests/`)

---

## Current Phase Snapshot

Phase: 2A (Go API)

Resume point:

- API routes are wired in `internal/handlers/tasks.go`
- In-memory store logic is in `internal/store/task_store.go`
- Route registration is in `cmd/main.go`
- Manual API verification is done via Bruno requests in `bruno/requests/`

Quick resume checklist for a new chat:

1. Confirm server runs: `go run ./cmd/main.go`
2. Confirm compile/test baseline: `go test ./...`
3. Run Bruno flow: create -> complete -> delete -> get
4. Add tests:
	- store unit tests in `internal/store/task_store_test.go`
	- handler tests in `internal/handlers/tasks_test.go` with `net/http/httptest`
5. Continue with SQLite persistence while keeping existing endpoints unchanged

---

## Project Structure

```
go-api/
├── cmd/
│   └── main.go           # Entry point, route registration
├── bruno/
│   ├── bruno.json
│   ├── environments/
│   │   └── local.bru
│   └── requests/
│       ├── 01-health.bru
│       ├── 02-get-tasks.bru
│       ├── 03-create-task.bru
│       ├── 04-delete-task.bru
│       └── 05-complete-task.bru
├── internal/
│   ├── handlers/
│   │   └── tasks.go      # HTTP handler functions
│   ├── models/
│   │   └── task.go       # Task struct
│   └── store/
│       └── task_store.go # In-memory task repository
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

## API Testing (Bruno)

Collection path: `go-api/bruno`

Request order for a full flow:

1. `01-health`
2. `02-get-tasks`
3. `03-create-task`
4. `05-complete-task`
5. `04-delete-task`
6. `02-get-tasks`

Use `params:path { id: ... }` in the delete/complete requests to target the task ID you want.

---

## Learning Goals

- Go module system and package layout
- `net/http` standard library (no framework)
- JSON marshaling/unmarshaling (`encoding/json`)
- HTTP method routing (`r.Method`)
- Path parameter parsing with `strings.Split` + `strconv.Atoi`
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

---

## Next Step

Move from in-memory store to SQLite-backed persistence while keeping the same HTTP API contract.
