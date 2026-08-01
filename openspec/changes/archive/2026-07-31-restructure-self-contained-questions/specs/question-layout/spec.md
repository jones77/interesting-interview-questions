# Question Layout

## ADDED Requirements

### Requirement: Canonical problem directory structure

Every question SHALL live in a directory
`questions/<difficulty>/<short-name>/` containing a `README.md` with the
problem statement and one subdirectory per supported language (`go/` now,
`python/` in a follow-up). Language directories SHALL be self-contained:
all code and tests for a problem reside inside them and SHALL NOT import or
reference files from other problems or from repo-internal packages.

#### Scenario: Structure of a question folder

- **WHEN** a question `two-sum` of difficulty `easy` exists
- **THEN** its files live under `questions/easy/two-sum/` with
  `README.md` and a `go/` subdirectory holding the Go module

#### Scenario: No cross-problem references

- **WHEN** any code or test file inside a question directory is inspected
- **THEN** it references only files within its own directory, the standard
  library, and (for generated stubs) nothing else

### Requirement: Problem statement in README.md only

The full problem statement SHALL appear in the generated `README.md` of a
question directory. Code and test files SHALL NOT duplicate the statement;
generated stubs SHALL contain at most a one-line doc comment pointing at
the `README.md`.

#### Scenario: Statement exists exactly once per language

- **WHEN** a question supports one or more languages
- **THEN** the problem statement text exists only in `README.md` and in
  no code or test file under the language subdirectories

#### Scenario: Stub points at the README

- **WHEN** a generated Go stub for question `two-sum` is inspected
- **THEN** its doc comment reads `TwoSum solves Two Sum. See README.md for
  the problem statement.` and contains no statement text

### Requirement: Self-contained Go programs

Each Go problem SHALL be its own Go module inside `questions/<difficulty>/
<short-name>/go/` with a `go.mod` declaring a bare module name
(`module <short-name-without-hyphens>`) and a `go` version matching the
root module. The package SHALL be `package main` with a trivial
`func main() {}`, so `go run .` compiles and exits successfully, and
`go test .` runs the tests in the same package.

#### Scenario: Program runs standalone

- **WHEN** a user runs `go run .` inside `questions/easy/two-sum/go/`
- **THEN** the program compiles and exits with status 0

#### Scenario: Tests run in the same package

- **WHEN** a user runs `go test .` inside `questions/easy/two-sum/go/`
- **THEN** the hand-written tests in `two-sum_test.go` execute without
  importing any other package

### Requirement: Inlined shared data types

Problems whose signatures use `TreeNode` or `ListNode` SHALL define those
types inside their own Go file; they SHALL NOT import a shared types
package.

#### Scenario: Tree problem defines its own node type

- **WHEN** a tree-based problem such as `invert-binary-tree` is generated
- **THEN** its stub file contains a local `TreeNode` definition and imports
  no shared type package
