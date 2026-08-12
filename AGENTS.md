# AGENTS.md

## Structure

Each question lives in `questions/<difficulty>/<short-name>/` with:
a generated `README.md` (the problem statement — a single copy, not
duplicated in code), and one self-contained language subdirectory per
implementation (`go/` and `python/`).

A Go problem is its own module: `questions/<difficulty>/<short-name>/go/`
holds a generated `go.mod`, a generated `package main` stub with a trivial
`func main() {}`, and a hand-written test file (`<short-name>_test.go`,
also `package main`) with the real tests that verify a solved
implementation.

A Python problem lives in `questions/<difficulty>/<short-name>/python/`
and holds a generated `<short_name>.py` stub and a hand-written
`test_<short_name>.py` file using the `unittest` module.

Problems never import repo-internal packages; shared `TreeNode`/`ListNode`
types are inlined per problem. Stubs must compile and must fail their tests
(the committed state is always unimplemented).

`questions.yaml` is the single source of truth for titles, difficulties,
categories, frequency, Go function signatures, Python function signatures,
and reference solutions.
**Never write a solution into a `.go` file** — only put stubs there.
If asked for a solution, point to the `solution` field in `questions.yaml`.

Questions and solutions are drawn from the model's training data, which may
include content from public coding platforms (LeetCode, HackerRank, etc.)
as well as other sources. They are not copied from a single site.

## Adding a Question

1. Add entry to `questions.yaml` (include `signature` for function-style
   questions or `kind` for struct-style — see existing entries). Also add
   `python_signature` for the generated Python stub.
2. Add a `solution` field to the entry in `questions.yaml`.
3. Write a Go test file at
   `questions/<difficulty>/<short-name>/go/<short-name>_test.go`
   (`package main`, same folder as the stub).
4. Write a Python test file at
   `questions/<difficulty>/<short-name>/python/test_<short_name>.py`
   (uses `unittest`).
5. Run `./bin/sync` to generate the `README.md`, `go.mod`, Go stub, and
   Python stub.
6. Commit `questions.yaml` and the new directory.

`short_name` must be unique — it becomes the directory name.

## Working with Solutions

The committed state of every problem is an unimplemented stub. If you
encounter a filled-in implementation while adding or regenerating problems,
stop and ask the user what they want to do. Solutions should be worked on in
a clean repo state separate from adding or recreating problems, because
regenerating stubs will overwrite implemented code.

## Commands

| Action             | Command                                  |
|--------------------|------------------------------------------|
| Build tooling      | `go build ./...` (root module only)      |
| Test one (Go)      | `cd questions/easy/two-sum/go && go test ./...` |
| Test one (Python)  | `cd questions/easy/two-sum/python && python3 -m unittest` |
| Test all questions | `./bin/test`                             |
| List / query       | `go run ./cmd/report/`                   |
| Flag docs          | `go run ./cmd/report/ -h`                |
| Regenerate + verify| `./bin/sync`                             |
| Force-regenerate   | `FORCE=1 ./bin/sync`                     |

`./bin/sync` regenerates stubs/READMEs, aborts on dirty generated files,
builds the root module, and verifies every problem builds and fails its
tests. `./bin/test` runs every problem's Go and Python test suites. The
root `go test ./...` does not reach problems — each has its own nested
module.

## Style

`.editorconfig` enforces 120-char lines,
2-space indent for YAML/Markdown,
tabs for Go files (gofmt standard),
4-space indent for Python files (PEP 8).
Use [semantic line breaks](https://sembr.org).
