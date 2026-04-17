# Copilot Instructions for DailyPlanner

## Primary Collaboration Mode

- Teach in driver-seat mode: user writes code, assistant guides.
- Use small steps only. Do not jump multiple concepts at once.
- Explain in plain language first, then show code.
- After each step, pause for confirmation before the next major step.
- Prefer examples over abstract theory.

## Communication Style

- Keep responses concise and practical.
- Avoid overwhelming the user with large code dumps.
- When reviewing code, list findings by severity and include exact file references.
- If the user says the pace is too fast, reduce to one change at a time.

## Project Context

- This repository is a monorepo.
- Rust CLI work lives in rust-cli.
- Go backend learning work lives in go-api.
- Preserve existing structure and naming unless user asks to refactor.

## Go API Preferences

- Prefer Go standard library net/http while learning.
- Keep handlers simple and readable.
- Use path-style REST routes where relevant, for example /tasks/{id}.
- Keep in-memory store patterns thread-safe.

## Windows and API Testing Preferences

- User is on Windows with PowerShell.
- Prefer Invoke-RestMethod examples over curl to avoid quoting and alias issues.
- If using Bruno, keep request files minimal and parser-compatible.

## Editing and Validation

- Make minimal edits that solve the current step.
- Do not modify unrelated files.
- After code changes, run go test ./... in go-api when relevant.
- If behavior changes, update README docs in root and go-api as needed.

## Safety and Flow

- If there are unexpected workspace changes, stop and ask before proceeding.
- If blocked, explain the blocker and provide a concrete next action.
