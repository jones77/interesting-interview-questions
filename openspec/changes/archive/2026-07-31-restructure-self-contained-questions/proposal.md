# Restructure Self-Contained Questions

## Why

Each interview question currently lives as a library package inside one
repo-wide Go module, sharing `internal/types`, `internal/testutil`, and
the root `go.mod`. This makes every question dependent on the whole repo:
you cannot `go run` or `go test` a single problem in isolation, and the
question text is duplicated into each `.go` doc comment.

We want the repo to become **multi-language** (Go now, Python next), with
each problem a **self-contained, runnable program** that compiles but fails
its tests — so users can clone, open a single folder, and work entirely
within it.

## What Changes

- **BREAKING** Restructure `questions/` to
  `questions/<difficulty>/<problem-name>/` with:
  - a generated `README.md` holding the problem statement (single copy,
    no duplication across languages), and
  - a `go/` subdirectory containing a self-contained Go program.
- **BREAKING** Each Go problem becomes its own Go module (`go.mod`) with a
  `package main` stub plus a trivial `func main()`, runnable via
  `go run .` and testable via `go test .` inside its folder.
- **BREAKING** Go stubs stop importing `internal/types` and
  `internal/testutil`; shared `TreeNode`/`ListNode` types are inlined per
  problem, and tests use stdlib assertions. Stubs contain no problem text —
  only a one-line pointer to the sibling `README.md`.
- `cmd/generate` grows from "write Go stubs" into "write README.md +
  self-contained Go stubs" for every question, still driven by
  `questions.yaml`.
- `bin/sync` reworks verification: regenerate stubs (dirty-stub detection
  for both `.go` files and `README.md`), then for **each** problem assert
  `go build ./...` succeeds and `go test ./...` fails (committed state is
  always failing stubs).
- Remove now-dead tooling: `internal/types`, `internal/testutil`, and
  `cmd/fix_tests` (its import-rewriting purpose disappears).
- Migrate all 40 existing questions to the new layout, including a
  mechanical rewrite of the 40 hand-written test files to be self-contained.
- Update user and maintainer docs (`README.md`) for the new per-problem
  workflow.

## Capabilities

### New Capabilities

- `question-layout`: the canonical `questions/` directory structure —
  problem statement in `README.md`, one language subdirectory per
  implementation (`go/` now, `python/` later), all files self-contained.
- `question-generation`: `cmd/generate` produces `README.md` and
  self-contained Go stub programs from `questions.yaml` (package main,
  own `go.mod`, inlined shared types, no problem text in code).
- `question-verification`: `bin/sync` verifies the committed state —
  stubs are pristine (unimplemented), every Go problem builds, and every
  problem's tests fail until solved.

### Modified Capabilities

None — no existing specs are present (`openspec/specs/` is empty).

## Impact

- **Repo structure**: all 40 `questions/<difficulty>/<name>/` folders move
  to the new layout; root module shrinks to `cmd/` + `internal/` tooling.
- **Go module graph**: nested per-problem `go.mod` files are automatically
  excluded from root `go build ./...`/`go test ./...`.
- **Tooling**: `cmd/generate` (extended), `bin/sync` (reworked),
  `cmd/report` (path/format tweaks), `cmd/fix_tests` + `internal/types` +
  `internal/testutil` (deleted).
- **Tests**: all hand-written test files rewritten mechanically once to be
  self-contained; their assertions and cases are preserved.
- **Docs**: `README.md` user/maintainer instructions change.
- **No changes** to `questions.yaml` schema in this change (Python fields
  arrive with the follow-up Python proposal).
