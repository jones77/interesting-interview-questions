# AGENTS.md

## Structure

Each interview question lives in a single `<short-name>.md` file under `questions/` —
e.g. `questions/two-sum.md`, `questions/fizzbuzz.md`.

Questions and their solutions are stored in separate YAML files.
The report tool generates `.md` files from the questions only and leaves the solutions out.

A **hint** is fine for problems where a hint is traditional.
The hint should point at the trouble spot without showing the way through it.

## Adding a Question

1. Add an entry to `questions.yaml`.
2. Add its solution to `solutions.yaml`.
3. Run `go run ./cmd/report/` to regenerate `.md` files.
4. Commit `questions.yaml`, `solutions.yaml`, and the new `.md` file.

`short_name` must be unique — it becomes the filename
(e.g. `two-sum` → `questions/two-sum.md`).

## Difficulty

Valid values: `easy`, `medium`, `hard`, `expert`. Every question needs one.
Assign based on well-known ratings (LeetCode, etc.).

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

Files should follow `.editorconfig` with a 120-character line limit.
Use [semantic line breaks](https://sembr.org) and wrap at clause boundaries.
