# AGENTS.md

## Structure

Each interview question lives in a single `<short-name>.md` file at repo root — e.g. `two-sum.md`, `fizzbuzz.md`.

Questions and their solutions are stored in `interview-questions.db` (SQLite). The `.md` files are generated output containing only the question — **never include the solution** in a `.md` file.

A **rot13 hint** is permissible if hints are traditionally given for that problem.

## Adding a Question

1. Add a new `Question` entry to `seedData()` in `cmd/populate/main.go`.
2. Run `go run ./cmd/populate/` to rebuild the DB.
3. Run `go run ./cmd/report/` to regenerate `.md` files.
4. Commit `cmd/populate/main.go`, `interview-questions.db`, and the new `.md` file.

## Developer Commands

| Action             | Command                     |
|--------------------|-----------------------------|
| Populate DB        | `go run ./cmd/populate/`    |
| Generate reports   | `go run ./cmd/report/`      |
| Build all          | `go build ./...`            |
| Test all           | `go test ./...`             |

## DB Schema

```
questions (id, short_name, title, question_text, solution, hint, difficulty, category)
```

- `question_text` → what goes in the `.md` file (the problem statement).
- `solution` → the answer, stored in the DB only, never written to markdown.
- `hint` → optional rot13-ciphered hint string.
