# AGENTS.md

## Structure

Each question lives in `questions/<difficulty>/<short-name>/` with:
a generated `README.md` (the problem statement — a single copy, not
duplicated in code), and one self-contained language subdirectory per
implementation (`go/` now, `python/` in a follow-up).

A Go problem is its own module: `questions/<difficulty>/<short-name>/go/`
holds a generated `go.mod`, a generated `package main` stub with a trivial
`func main() {}`, and a hand-written test file (`<short-name>_test.go`,
also `package main`) with the real tests that verify a solved
implementation. Problems never import repo-internal packages; shared
`TreeNode`/`ListNode` types are inlined per problem. Stubs must compile
and must fail their tests (the committed state is always unimplemented).

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
3. Write a test file at
   `questions/<difficulty>/<short-name>/go/<short-name>_test.go`
   (`package main`, same folder as the stub).
4. Run `./bin/sync` to generate the `README.md`, `go.mod`, and stub.
5. Commit `questions.yaml` and the new directory.

`short_name` must be unique — it becomes the directory name.

## Commands

| Action             | Command                                  |
|--------------------|------------------------------------------|
| Build tooling      | `go build ./...` (root module only)      |
| Test one           | `cd questions/easy/two-sum/go && go test ./...` |
| Test all questions | `./bin/test`                             |
| List / query       | `go run ./cmd/report/`                   |
| Flag docs          | `go run ./cmd/report/ -h`                |
| Regenerate + verify| `./bin/sync`                             |
| Force-regenerate   | `FORCE=1 ./bin/sync`                     |

`./bin/sync` regenerates stubs/READMEs, aborts on dirty generated files,
builds the root module, and verifies every problem builds and fails its
tests. `./bin/test` runs every problem's test suite. The root `go test
./...` does not reach problems — each has its own nested module.

## Style

`.editorconfig` enforces 120-char lines,
2-space indent for YAML/Markdown,
tabs for Go files (gofmt standard).
Use [semantic line breaks](https://sembr.org).
