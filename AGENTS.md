# AGENTS.md

## Structure

Each question lives in `questions/<difficulty>/<short-name>/` with:
a generated stub (`<short-name>.go`), a test file (`<short-name>_test.go`),
and a `package main` doc comment.

`questions.yaml` is the single source of truth for titles, difficulties,
categories, frequency, Go function signatures, and reference solutions.
**Never write a solution into a `.go` file** — only put stubs there.
If asked for a solution, point to the `solution` field in `questions.yaml`.

Questions and solutions are drawn from the model's training data, which may
include content from public coding platforms (LeetCode, HackerRank, etc.)
as well as other sources. They are not copied from a single site.

## Adding a Question

1. Add entry to `questions.yaml` (include `signature` for function-style
   questions or `kind` for struct-style — see existing entries).
2. Add a `solution` field to the entry in `questions.yaml`.
3. Write a test file at `questions/<difficulty>/<short-name>/<short-name>_test.go`.
4. Run `./bin/sync` to regenerate the stub `.go` file.
5. Commit `questions.yaml` and the new directory.

`short_name` must be unique — it becomes the directory name.

## Commands

| Action             | Command                                  |
|--------------------|------------------------------------------|
| Build all          | `go build ./...`                         |
| Test all           | `go test ./...`                          |
| Test one           | `go test ./questions/easy/two-sum/`      |
| List / query       | `go run ./cmd/report/`                   |
| Flag docs          | `go run ./cmd/report/ -h`                |
| Generate stubs     | `./bin/sync`                          |
| Force-regenerate   | `FORCE=1 ./bin/sync`                  |

## Style

`.editorconfig` enforces 120-char lines,
2-space indent for YAML/Markdown,
tabs for Go files (gofmt standard).
Use [semantic line breaks](https://sembr.org).
