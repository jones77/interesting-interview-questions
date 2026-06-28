# AGENTS.md

## Structure

Each interview question lives in a single `<short-name>.md` file under `questions/` — e.g. `questions/two-sum.md`, `questions/fizzbuzz.md`.

Questions and their solutions are stored as YAML. The `.md` files are generated output containing only the question — **never include the solution** in a `.md` file.

A **rot13 hint** is permissible if hints are traditionally given for that problem.

## Adding a Question

1. Add an entry to `questions.yaml`.
2. Add its solution to the decoded `solutions-plain.yaml`, then run `./bin/rot13 solutions-plain.yaml > solutions.yaml` to re-encode.
3. Run `go run ./cmd/report/` to regenerate `.md` files.
4. Commit `questions.yaml`, `solutions.yaml`, and the new `.md` file.

`short_name` must be unique — it becomes the filename (e.g. `two-sum` → `questions/two-sum.md`).

## Developer Commands

| Action             | Command                     |
|--------------------|-----------------------------|
| Generate reports   | `go run ./cmd/report/`      |
| Build all          | `go build ./...`            |
| Test all           | `go test ./...`             |

## Viewing Solutions

`solutions.yaml` is ROT13-encoded. Decode it with `./bin/rot13`:

```sh
./bin/rot13 solutions.yaml
```

## Data Files

- `questions.yaml` — question metadata and text (plaintext, committed).
- `solutions.yaml` — ROT13-encoded short_name → solution map (committed).
- `solutions-plain.yaml` — decoded working copy (gitignored, recreate via `./bin/rot13 solutions.yaml > solutions-plain.yaml`).
