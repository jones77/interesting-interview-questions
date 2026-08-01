# Question Verification

## Purpose

Specify how the repo verifies generated questions stay pristine: stub and
README regeneration checks, per-problem builds, the failing-test invariant
for committed stubs, and an all-question test runner.

## Requirements

### Requirement: Pristine stub enforcement

`bin/sync` SHALL detect stubs that have been modified away from the
generated output (i.e., implemented by a user) before doing anything else.
If any stub differs from the freshly generated version, `bin/sync` SHALL
restore the modified files, report them, and abort with a non-zero exit
status unless `FORCE=1` is set, in which case it SHALL overwrite them.

#### Scenario: Implemented stub aborts sync

- **WHEN** a user has implemented the body of a stub and runs `./bin/sync`
- **THEN** sync restores the generated stub, prints the offending files,
  and exits non-zero

#### Scenario: Force regeneration

- **WHEN** a maintainer runs `FORCE=1 ./bin/sync` after changing
  `questions.yaml` signatures
- **THEN** sync overwrites all stubs with freshly generated content and
  completes successfully

### Requirement: README freshness

`bin/sync` SHALL regenerate `README.md` files alongside stubs and SHALL
treat any difference from the committed version as a dirty state that
aborts the run (or is overwritten under `FORCE=1`).

#### Scenario: Stale README is reported

- **WHEN** a committed `README.md` differs from the generated content and
  `FORCE` is unset
- **THEN** sync reports the file and aborts non-zero

### Requirement: Per-problem build

For every question, `bin/sync` SHALL run `go build ./...` inside
`questions/<difficulty>/<short-name>/go/` and require it to succeed. It
SHALL also build the root module (`cmd/`, `internal/`) successfully.

#### Scenario: Broken problem fails sync

- **WHEN** a question's Go module fails to compile
- **THEN** sync reports the failing directory and exits non-zero

### Requirement: Failing-test invariant for committed stubs

In the committed state (pristine stubs), `bin/sync` SHALL run
`go test ./...` inside every question's `go/` directory and require that
it **fails** (non-zero exit). A passing test run indicates the stub has
been implemented, which is not a valid committed state for this repo.

#### Scenario: Stub fails its tests

- **WHEN** all stubs are pristine and `./bin/sync` runs
- **THEN** every question's `go test ./...` exits non-zero and sync
  completes successfully

#### Scenario: Passing tests are flagged

- **WHEN** any question's `go test ./...` exits zero (stub passes)
- **THEN** sync reports the directory as implemented and exits non-zero

### Requirement: All-question test runner

The repo SHALL provide a `bin/test` command that walks every language
directory under `questions/` and runs its test suite, so users can run
all tests across all problems with one command.

#### Scenario: Running every test

- **WHEN** a user runs `./bin/test`
- **THEN** the test suite of every question directory is executed
