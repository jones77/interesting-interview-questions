# AGENTS.md

## Structure

Each interview question lives in a single `<short-name>.md` file under `questions/` —
e.g. `questions/two-sum.md`, `questions/fizzbuzz.md`.

Questions and their solutions are stored as YAML.
The `.md` files are generated output containing only the question —
**never include the solution** in a `.md` file.

A **hint** is permissible if hints are traditionally given for that problem.
A hint must not give the game away — it should point at the difficulty
without revealing the approach.

## Adding a Question

1. Add an entry to `questions.yaml`.
2. Add its solution to `solutions.yaml`.
3. Run `go run ./cmd/report/` to regenerate `.md` files.
4. Commit `questions.yaml`, `solutions.yaml`, and the new `.md` file.

`short_name` must be unique — it becomes the filename
(e.g. `two-sum` → `questions/two-sum.md`).

## Difficulty

Valid values: `easy`, `medium`, `hard`, `expert`. Every question must have one.

## Developer Commands

| Action             | Command                     |
|--------------------|-----------------------------|
| Generate reports   | `go run ./cmd/report/`      |
| Build all          | `go build ./...`            |
| Test all           | `go test ./...`             |

## Data Files

- `questions.yaml` — question metadata and text (plaintext, committed).
- `solutions.yaml` — short_name to solution map (plaintext, committed).

## Formatting

Files should follow `.editorconfig` (`max_line_length = 120`).
