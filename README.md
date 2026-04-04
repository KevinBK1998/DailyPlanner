# 📋 Daily Planner - Full-Stack Learning Project

A comprehensive multi-language, multi-platform todo ecosystem. Learn Rust, Go, Java, React, TypeScript, Kotlin, and more through one coherent project that spans CLI, REST API, web frontend, and mobile app.

**Phases:** CLI (Rust) → Backend (Go/Spring Boot) → Frontend (React/TS) → Mobile (Android/Kotlin) → Integration

> **Note:** This project structure and roadmap were generated with AI assistance. The code you'll write during each phase is 100% yours!

---

## 🚀 Learning Roadmap

### Phase 1: Foundations (CLI + Core Concepts) ✅ **Completed**

### Phase 2: Backend Development (Go/Spring) ✨ **Current Phase**

---

## 🎯 Workspace Getting Started

This repository is organized as a multi-language monorepo.

- Rust CLI: [rust-cli/README.md](rust-cli/README.md)
- Go REST API: [go-api/README.md](go-api/README.md)
- Backend phases (Spring): `java-api/`
- Frontend/mobile phases: `react-web/` and `android-app/`

### Rust quick commands
```bash
cd rust-cli
cargo build
cargo run
cargo test
```

### Go quick commands
```bash
cd go-api
go run ./cmd/main.go
go test ./...
```

---

## 📚 Project Structure

```
daily_planner/
├── rust-cli/
│   ├── Cargo.toml
│   ├── README.md
│   ├── src/
│   │   ├── main.rs
│   │   ├── cli.rs
│   │   ├── manager.rs
│   │   └── models.rs
│   └── data/
├── .github/
│   └── workflows/
│       ├── rust-cli-ci.yml
│       └── rust-cli-release.yml
├── go-api/             # Phase 2A ✨ current
├── java-api/           # Phase 2B target
├── react-web/          # Phase 3 target
├── android-app/        # Phase 4 target
└── README.md
```

---

## 🧩 Subproject Docs

- Rust CLI details (architecture, commands, test/release): [rust-cli/README.md](rust-cli/README.md)
- Go REST API details (structure, endpoints, learning goals): [go-api/README.md](go-api/README.md)

---

## 🚢 Release Workflows

- CI for Rust CLI: [.github/workflows/rust-cli-ci.yml](.github/workflows/rust-cli-ci.yml)
- Tagged binary release: [.github/workflows/rust-cli-release.yml](.github/workflows/rust-cli-release.yml)

Release tag format for Rust CLI binaries:

- `rust-cli-v*` (example: `rust-cli-v0.1.0`)

---

## 🔮 Full Learning Roadmap

### Phase 2: Backend Development

#### 2A - Go REST API
Build a lightweight backend to expose todos via HTTP endpoints.

See [go-api/README.md](go-api/README.md) for full details, endpoints, and learning goals.

**Tech:** Go, standard library (net/http), SQLite

---

#### 2B - Java Spring Boot (Alternative)
Enterprise-grade backend with proven patterns.

**Same Endpoints as Go API**

**Learning Goals:**
- Object-oriented design
- Dependency injection (Spring beans)
- MVC architecture
- ORM with JPA/Hibernate
- Spring Boot conventions

**Tech:** Java, Spring Boot, Spring Data JPA, H2/PostgreSQL

**Goal:** Experience both lightweight (Go) and enterprise (Java) backend styles

---

### Phase 3: Frontend Development

#### 3A - HTML + CSS Foundation
Build a clean, responsive static interface.

**Features:**
- Task list display
- Add task form
- Responsive design (mobile-first)
- Custom theme/styling

**Learning Goals:**
- Semantic HTML5
- CSS Grid & Flexbox
- Responsive design patterns
- Accessibility (a11y)

**Tech:** HTML5, CSS3, no frameworks

---

#### 3B - JavaScript/TypeScript + React
Create an interactive, type-safe frontend.

**Components:**
- `TaskList` - Display tasks
- `AddTaskForm` - Add new tasks
- `TaskItem` - Individual task management
- State management with hooks

**Learning Goals:**
- React component architecture
- Hooks (`useState`, `useReducer`, custom hooks)
- Async API calls (fetch/axios)
- Type safety with TypeScript interfaces
- State management patterns
- Testing components

**Tech:** React, TypeScript, Vite, Axios/Fetch API

**Goal:** Master modern frontend engineering with type safety

---

### Phase 4: Mobile Development

#### 4A - Android App (Java/Kotlin)
Bring the todo manager to mobile devices.

**Features:**
- View todos in native Android UI
- Add/edit/delete tasks
- Local persistence (Room DB)
- Backend sync (optional)
- Push notifications

**Learning Goals:**
- Activity lifecycle & fragments
- Material Design UI components
- Room database (local persistence)
- SharedPreferences for app settings
- HTTP requests on mobile (Retrofit/OkHttp)
- Background tasks & notifications

**Tech:** Kotlin, Android SDK, Room, Jetpack libraries

**Goal:** Port the todo app to mobile with native Android patterns

---

### Phase 5: Integration & Advanced Features

#### 5A - Cross-Platform Synchronization
Connect all components together.

**Integration Points:**
- Rust CLI ↔ Go/Java backend (API calls)
- React frontend ↔ Backend API
- Android app ↔ Backend API
- Real-time sync across devices

**Tech:** REST API calls, webhook listeners, sync algorithms

---

#### 5B - Advanced Todo Features

**New Features:**
- **Categories** - Organize todos by project/area
- **Priorities** - High/Medium/Low priority levels
- **Deadlines** - Due dates and reminders
- **Recurring Tasks** - Daily/weekly/monthly patterns
- **Notifications** - Cron jobs, push notifications, email reminders
- **Gamification** - XP system, streak tracking, achievement badges

**Learning Goals:**
- Database schema design for complex features
- Scheduling (cron jobs, task queues)
- Push notification systems
- Caching strategies
- Real-time updates (WebSockets)

**Tech:** Task queues (Bull, Celery), Redis for caching, Firebase Cloud Messaging or similar

**Goal:** Build a fully-featured, engaging todo ecosystem across all platforms

---

## 🎯 Complete Architecture Overview

```
┌─────────────────────────────────────────────────────┐
│          Phase 1: CLI (Rust)                        │
│     (Core logic & local file management)            │
└────────────┬────────────────────────────────────────┘
             │
             ├─────────────────────────────────────┐
             │                                     │
┌────────────▼──────────────┐      ┌───────────────▼────────────┐
│  Phase 2A: Go API         │      │  Phase 2B: Spring Boot     │
│  (Lightweight backend)    │      │  (Enterprise backend)      │
├───────────────────────────┤      ├────────────────────────────┤
│ - REST endpoints          │      │ - REST endpoints           │
│ - Goroutines              │      │ - Spring beans             │
│ - Concurrency patterns    │      │ - JPA/ORM                  │
│ - SQLite/Postgres         │      │ - Dependency injection     │
└────────────┬──────────────┘      └────────────┬───────────────┘
             │                                  │
             └──────────┬───────────────────────┘
                        │
        ┌───────────────┴──────────────┐
        │                              │
   ┌────▼──────────────┐    ┌──────────▼─────────┐
   │ Phase 3: Frontend │    │ Phase 4: Android   │
   ├───────────────────┤    ├────────────────────┤
   │ - HTML/CSS        │    │ - Kotlin/Java      │
   │ - React + TS      │    │ - Material Design  │
   │ - Responsive UI   │    │ - Room DB          │
   └────┬──────────────┘    └────────┬───────────┘
        │                            │
        └────────────┬───────────────┘
                     │
         ┌───────────▼────────────┐
         │   Phase 5: Integration │
         ├────────────────────────┤
         │ - Cross-device sync    │
         │ - Advanced features    │
         │ - Gamification         │
         │ - Notifications        │
         │ - Full-stack ecosystem │
         └────────────────────────┘
```

---

## 📖 Learning Path Summary

1. **Phase 1 (Rust CLI)** → Master fundamentals: ownership, borrowing, structs, enums, file I/O
2. **Phase 2A (Go)** → Learn concurrency & lightweight API design
3. **Phase 2B (Spring Boot)** → Master enterprise patterns & OOP
4. **Phase 3A (HTML/CSS)** → Build accessible, responsive interfaces
5. **Phase 3B (React/TypeScript)** → Modern component-based frontend with type safety
6. **Phase 4 (Android/Kotlin)** → Native mobile development
7. **Phase 5 (Integration)** → Cross-platform sync & advanced features

**By the end:** A full-stack, cross-platform todo ecosystem spanning CLI, web, and mobile! 🚀

---

## 🎯 Suggested Learning Order

Follow this sequence to build confidence progressively:

1. **Rust CLI** → Establish low-level programming fundamentals and memory safety
2. **Go Backend** → Learn APIs, HTTP routing, and concurrency patterns
3. **React + TypeScript Frontend** → Master modern web development with type safety
4. **Java Spring Boot OR Android App** → Choose based on interest:
   - Spring Boot for enterprise backend skills
   - Android for mobile platform experience
5. **Integration Phase** → Connect everything together and add advanced features

**Why this order?**
- Start simple (CLI) → build confidence
- Add HTTP layer (Go) → understand distributed systems
- Add UI layer (React) → see user-facing results
- Add mobile/enterprise (Android/Spring) → expand horizons
- Integrate (Phase 5) → synthesize everything

---

## ✨ Bonus: Make It Fun!

### 🎨 Pokémon-Inspired Themes
Bring personality to your todo app with themed UI variants:

- **Umbreon Theme** - Dark mode with sleek purple accents (perfect for late-night coding)
- **Greninja Theme** - Sleek water/blue aesthetic with smooth animations
- **Charizard Theme** - Warm orange/fire tones for high-energy productivity
- **Dragonite Theme** - Professional gold/blue for serious work sessions

### 💎 Gamification Features
Make task management addictive:

- **XP System** - Earn experience points for completing tasks
- **Streak Tracking** - Maintain daily/weekly task completion streaks
- **Achievement Badges** - Unlock badges (1st task, 7-day streak, 100 tasks, etc.)
- **Level Progression** - Unlock new features as you level up
- **Leaderboards** - (optional) Compare streaks with friends

### 🎯 UI/UX Inspired by Modern Chat Apps
- Clean, minimal interface (like Discord/Slack)
- Invisible-but-readable color schemes
- Smooth animations and transitions
- Task cards with subtle hover effects
- Dark mode as primary theme

### 📱 Cross-Platform Consistency
Same UI/UX across:
- CLI (ASCII art, colored text)
- Web (React with chosen theme)
- Android (Material Design variant of theme)

---

## 🔨 Current Phase (Phase 1: Rust CLI)

### Prerequisites
- [Rust installed](https://rustup.rs/)

### Building
```bash
cargo build
```

### Running
```bash
cargo run
```

### Cleaning
```bash
cargo clean
```

---

## 📚 Project Structure

```
daily_planner/
├── Cargo.toml          # Project manifest
├── src/
│   └── main.rs         # Application code
└── README.md           # This file
```

---

## 🎓 What You'll Learn in Phase 1

This foundational project covers essential Rust concepts through hands-on implementation:

| Concept | Why It Matters |
|---------|----------------|
| **Ownership** | Rust's memory safety guarantee without a garbage collector |
| **Borrowing** | Safe multi-reference patterns (& for immutable, &mut for mutable) |
| **Structs** | Organizing related data together |
| **Enums** | Type-safe alternatives to magic strings/numbers |
| **Pattern Matching** | Safe, exhaustive conditional logic |
| **File I/O** | Reading/writing persistent data |
| **Error Handling** | Using `Result` and the `?` operator |
| **Collections** | Managing groups of data with `Vec` |