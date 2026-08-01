# Tasks: Self-Contained Questions Restructure

## 1. Generator rework

- [x] 1.1 Create `internal/generate/` package: load `questions.yaml`, iterate specs, route per-language emission
- [x] 1.2 Implement README renderer (`internal/generate/readme.go`): `# <Title>`, URL link when present, `question_text` verbatim
- [x] 1.3 Implement Go stub emitter (`internal/generate/go.go`) for function-style questions: `package main`, exported function from `signature`, zero-value returns, trivial `func main() {}`, one-line doc comment pointing at `README.md`
- [x] 1.4 Implement kind-style stub emitters for `lru-cache`, `codec-tree`, `max-path-tree` (struct/method stubs with zero-value bodies + `main()`)
- [x] 1.5 Inline `TreeNode`/`ListNode` definitions in stubs for `tree`/`list` kinds and kind-style questions that use them; drop `internal/types` imports
- [x] 1.6 Implement `go.mod` emitter: `module <short-name-without-hyphens>` with the Go version read from the root `go.mod`
- [x] 1.7 Rewire `cmd/generate` to call `internal/generate` and verify output for one function-style, one tree, and one kind-style question (build + run)

## 2. Verification rework

- [x] 2.1 Rework `bin/sync` Phase 1: regenerate stubs, `go.mod`, and `README.md` to a temp tree; diff against committed files; restore + abort non-zero when dirty unless `FORCE=1`
- [x] 2.2 Rework `bin/sync` Phase 2: `go build ./...` on the root module (tooling only)
- [x] 2.3 Add `bin/sync` per-problem checks: for every `questions/*/*/go/` run `go build ./...` (must succeed) and `go test ./...` (must exit non-zero)
- [x] 2.4 Add `bin/test`: walk all language directories under `questions/` and run their test suites
- [x] 2.5 Verify `bin/sync` aborts cleanly on an intentionally implemented stub and on a stale README

## 3. Migration of existing questions

- [x] 3.1 Build temporary `cmd/migrate`: move each of the 40 problems into `go/`, write `README.md` + `go.mod`, rewrite the test file (package `main`, strip package qualifiers, unqualify `types.*`, replace `testutil.*` with stdlib, drop `internal/*` imports)
- [x] 3.2 Dry-run migration on a representative sample (e.g., `two-sum`, `invert-binary-tree`, `lru-cache`, `serialize-deserialize-binary-tree`) and confirm each builds and fails tests
- [x] 3.3 Run migration across all 40 questions
- [x] 3.4 Fix per-problem fallout until `./bin/sync` passes end-to-end

## 4. Dead code removal and tooling cleanup

- [x] 4.1 Delete `cmd/fix_tests`, `internal/types`, `internal/testutil`
- [x] 4.2 Grep repo for references to deleted packages; fix stragglers; `go build ./...` on root passes
- [x] 4.3 Verify `go run ./cmd/report/` still lists questions correctly; adjust path/format references if any

## 5. Documentation

- [x] 5.1 Update root `README.md`: user flow (pick problem, read `README.md`, implement in `go/`, `go test .` there), maintainer flow, and the add-a-question steps
- [x] 5.2 Update `AGENTS.md` structure section and command table (`./bin/test`, per-problem `go test`)

## 6. Final gate

- [x] 6.1 From a clean tree: `./bin/sync` exits 0 (stubs pristine, all problems build, all tests fail)
- [x] 6.2 Spot-check `go run .`, `go test .`, and `go run ./cmd/report/` in at least one problem per difficulty
- [x] 6.3 Delete `cmd/migrate`; confirm nothing references it
