# Go REST API — Daily Planner

Phase 2A of the [DailyPlanner monorepo](../README.md). A lightweight REST API built with Go's standard library.

---

## Status

✅ **Milestone Complete (Phase 2A)**

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
- SQLite-backed task store (`internal/store/task_store.go`) after migrating from the in-memory map version
- Store and handler tests (`internal/store/task_store_test.go`, `internal/handlers/tasks_test.go`)
- Concurrent store tests for write-only and mixed read/write access
- Explicit server startup error handling and graceful shutdown in `cmd/main.go`
- Bruno request collection (`bruno/requests/`)

Test coverage highlights:

- Store: add/list, delete/list, complete, not-found error behavior, concurrent writes, and mixed read/write concurrency
- Handlers: create/list happy path, invalid JSON, invalid ID format, not-found IDs, unsupported methods, and unsupported paths

SQLite concurrency tuning:

- `journal_mode = WAL` to improve read/write overlap behavior
- `busy_timeout = 5000` to wait briefly on lock contention instead of failing fast
- `SetMaxOpenConns(1)` and `SetMaxIdleConns(1)` to align with SQLite's single-writer model and avoid `SQLITE_BUSY` spikes under concurrent tests

---

## Current Phase Snapshot

Phase: 2A (Go API)

Resume point:

- API routes are wired in `internal/handlers/tasks.go`
- SQLite store logic is in `internal/store/task_store.go` (in-memory map implementation was the previous step)
- Route registration is in `cmd/main.go`
- Manual API verification is done via Bruno requests in `bruno/requests/`

Quick resume checklist for a new chat:

1. Confirm server runs: `go run ./cmd/main.go`
2. Confirm compile/test baseline: `go test ./...`
3. Run Bruno flow: create -> complete -> delete -> get
4. Keep tests green:
	- store tests in `internal/store/task_store_test.go`
	- handler tests in `internal/handlers/tasks_test.go` with `net/http/httptest`
5. Continue with concurrency patterns and transaction-focused refactors when ready

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
│   │   ├── tasks.go      # HTTP handler functions
│   │   └── tasks_test.go # Handler tests
│   ├── models/
│   │   └── task.go       # Task struct
│   └── store/
│       ├── task_store.go      # SQLite task repository
│       └── task_store_test.go # Store tests
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
- Go maps and in-memory state (initial implementation)
- SQLite persistence
- API/store testing with `testing` and `net/http/httptest`
- Goroutines and concurrency patterns (covered with concurrent store tests)

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

CI workflow: [../.github/workflows/go-api-ci.yml](../.github/workflows/go-api-ci.yml)

It runs on push/PR changes for `go-api/**` and includes:

- `gofmt` check
- `go vet ./...`
- `go test ./...`

---

## Release

Release workflow: [../.github/workflows/go-api-release.yml](../.github/workflows/go-api-release.yml)

Release trigger tag format:

- `go-api-v*` (example: `go-api-v0.1.0`)

The workflow builds Linux and Windows binaries for `go-api/cmd/main.go` and publishes them to a GitHub Release.

---

## Next Step

Phase 2A goals are complete. Continue with either:

- Phase 2B (`java-api/`) for Spring Boot backend learning, or
- Phase 3 (`react-web/`) to build the frontend and integrate with this API.
