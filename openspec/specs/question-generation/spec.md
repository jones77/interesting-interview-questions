# Question Generation

## Purpose

Specify how generated artifacts — READMEs, `go.mod` files, and Go stubs —
are produced from `questions.yaml`, the single source of truth for
question content and reference solutions.

## Requirements

### Requirement: questions.yaml is the single source of truth

`cmd/generate` SHALL read `questions.yaml` and produce every generated
artifact for every question entry: the `README.md`, the `go.mod`, and the
Go stub. No generated file SHALL be committed that differs from what
`cmd/generate` produces for the same input.

#### Scenario: Generation covers all questions

- **WHEN** `cmd/generate` runs against the current `questions.yaml`
- **THEN** every entry yields a `README.md`, a `go/` directory with a
  `go.mod`, and a stub `.go` file at the canonical paths

#### Scenario: Generation is deterministic

- **WHEN** `cmd/generate` runs twice with unchanged `questions.yaml`
- **THEN** the two runs produce byte-identical files

### Requirement: README generation

The generated `README.md` SHALL contain the question title as a level-1
heading, a link to the source URL when present in `questions.yaml`, and
the full `question_text` verbatim.

#### Scenario: README contains the statement

- **WHEN** question `fizzbuzz` has `title: Fizz Buzz`, a URL, and a
  `question_text` field
- **THEN** `questions/easy/fizzbuzz/README.md` starts with
  `# Fizz Buzz`, links to the URL, and contains the `question_text`
  unchanged

### Requirement: Go stub generation

For each question, `cmd/generate` SHALL write a stub Go file in
`questions/<difficulty>/<short-name>/go/` that:

- declares `package main` plus a trivial `func main() {}`,
- exposes the signature from `questions.yaml` with the first letter
  capitalized,
- returns zero values for every return type,
- uses the kind-specific template for `lru-cache`, `codec-tree`, and
  `max-path-tree` questions (struct/method stubs with zero-value bodies),
- inlines `TreeNode` and `ListNode` definitions instead of importing a
  shared types package, and
- carries a one-line doc comment naming the function and pointing at the
  sibling `README.md`.

It SHALL also write the `go.mod` with module name
`<short-name-without-hyphens>` and the root module's Go version.

#### Scenario: Function-style stub

- **WHEN** `two-sum` has signature `twoSum(nums []int, target int) []int`
- **THEN** `questions/easy/two-sum/go/two-sum.go` contains
  `func TwoSum(nums []int, target int) []int` returning `nil`

#### Scenario: Kind-style stub

- **WHEN** `lru-cache` has `kind: lru-cache`
- **THEN** its stub file contains an `LRUCache` struct with `Constructor`,
  `Get` (returning `-1`), and `Put` (no-op) methods and a trivial `main()`

#### Scenario: Tree and list stubs inline their node types

- **WHEN** a question signature references `TreeNode` or `ListNode`
- **THEN** the stub file defines that type locally and imports no shared
  type package

### Requirement: No solutions in generated code

Generated stubs SHALL contain only signatures, zero-value returns, and
trivial bodies — never reference-solution logic. Reference solutions
remain exclusively in the `solution` field of `questions.yaml`.

#### Scenario: Stub has no solution logic

- **WHEN** the stub for any question is inspected
- **THEN** its function and method bodies return zero values only and
  contain no algorithm logic
