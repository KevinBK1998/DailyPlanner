# React Web - Daily Planner

Phase 3 of the [DailyPlanner monorepo](../README.md).

## Status

✨ In Progress (Phase 3A: HTML + CSS Foundation)

## Quick Start

From the react-web folder:

- Open index.html directly in your browser, or
- Use the VS Code Live Server extension to run a local static server.

## Dedicated Progress

### Current Milestone: 3A HTML + CSS Foundation

Completed:

- [x] Semantic page structure with header, form section, and task list section
- [x] Accessible form label/input setup for adding a task
- [x] Task rows with pending and completed visual states
- [x] Responsive layout updates for desktop and small-screen behavior
- [x] Keyboard focus visibility for text input, button, task rows, and checkboxes
- [x] Checkbox + label linkage for better accessibility and click targets
- [x] Scoped CSS selectors to avoid style bleed across unrelated elements
- [x] Long task title handling to prevent layout overflow

Current files:

- index.html
- styles.css

### Resume Point

If you are resuming in a new chat, continue from these checks:

1. Verify current static UI in browser (desktop and narrow mobile width).
2. Validate keyboard navigation order and focus rings.
3. Keep CSS scoped and avoid global selectors unless intentional.
4. Finish any remaining visual polish tasks for Phase 3A.
5. Start Phase 3B by adding JavaScript interactions, then migrate to React + TypeScript.

## Next Step

Complete final static polish for Phase 3A, then begin Phase 3B:

- Add task interactions (create, complete toggle, delete) in plain JavaScript first.
- Introduce React component structure after behavior is clear.
- Connect frontend actions to the Go API endpoints.
