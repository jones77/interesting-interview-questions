# AGENTS.md

## Structure

Each interview question lives under `questions/<difficulty>/<short-name>/` —
e.g. `questions/easy/two-sum/`. The test file and solution stub
all live in that directory. The solution stub's `package main` comment
contains the problem statement so the user sees the question when they open
the `.go` file.

Questions and their reference solutions are stored in separate YAML files.

A **hint** is fine for problems where a hint is traditional.
The hint should point at the trouble spot without showing the way through it.

**Never implement a solution for the user.** If they ask for one,
point them to `solutions.yaml` and tell them to solve it themselves first.

## Adding a Question

1. Add an entry to `questions.yaml`.
2. Add its solution to `solutions.yaml`.
3. Write a doc comment before the `package main` statement of the solution stub
   with the problem statement (title, description, and at least one example).
   Copy `templates/skeleton_solution.go.tpl` and `templates/skeleton_test.go.tpl`
   into the question directory and implement.
4. Commit `questions.yaml`, `solutions.yaml`, and the new question directory.

`short_name` must be unique — it becomes the directory name
(e.g. `two-sum` with `easy` → `questions/easy/two-sum/`).

## Difficulty

Valid values: `easy`, `medium`, `hard`, `expert`. Every question needs one.
Assign based on well-known ratings (LeetCode, etc.).

## Developer Commands

| Action             | Command                     |
|--------------------|-----------------------------|
| List questions     | `go run ./cmd/report/`      |
| Build all          | `go build ./...`            |
| Test all           | `go test ./...`             |
| Test one           | `go test ./questions/easy/two-sum/` |

Run `go run ./cmd/report/ -h` for full flag documentation. Flags `-q` (questions) and `-a` (all) output YAML; `-a` is equivalent to `-q -s`. `-l` or `--list` produces the same single-line output as no flags. `-s` requires `-f` and shows just the solution for filtered questions. `-f` alone shows just the question text. `-a` and `-f` are mutually exclusive.

## Data Files

- `questions.yaml` — question metadata and text (plaintext, committed).
- `solutions.yaml` — short_name to solution map (plaintext, committed).

## Formatting

Files should follow `.editorconfig` with a 120-character line limit.
Use [semantic line breaks](https://sembr.org) and wrap at clause boundaries.
