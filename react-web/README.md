# React Web - Daily Planner

Phase 3 of the [DailyPlanner monorepo](../README.md).

## Status

✨ In Progress (Phase 3B started: JavaScript interactions)

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

### Current Milestone: 3B JavaScript Foundation (Started)

Completed:

- [x] Added `app.js` and connected it to `index.html` with `defer`
- [x] Form submit handler prevents page reload
- [x] New task title is validated and appended to the DOM task list
- [x] Form reset + focus return after successful submit

Next in this milestone:

- [ ] Checkbox toggle updates row state (`pending` <-> `completed`)
- [ ] Delete action for task rows
- [ ] Keep IDs unique and stable for dynamic tasks

Current files:

- index.html
- styles.css
- app.js

### Resume Point

If you are resuming in a new chat, continue from these checks:

1. Verify current static UI in browser (desktop and narrow mobile width).
2. Validate keyboard navigation order and focus rings.
3. Submit a new task and confirm it appears in the task list.
4. Keep CSS scoped and avoid global selectors unless intentional.
5. Continue Phase 3B by wiring checkbox toggle behavior for dynamic and existing rows.

## Next Step

Continue Phase 3B with one small interaction step at a time:

- Add checkbox toggle behavior (`pending` -> `completed` and back).
- Add delete behavior for task rows.
- After core interactions are stable, move to React + TypeScript and API integration.
