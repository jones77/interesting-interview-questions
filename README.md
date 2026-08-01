# Interesting Interview Questions

Clone this repo, implement your own solutions, and make the tests pass.
Reference solutions are in `questions.yaml` for when you want to check
your work — solve it first, then peek.

Each problem is a self-contained program: the problem statement lives in
the problem's `README.md`, and each supported language has its own folder
(`go/` now, `python/` in the future) with code and tests that need nothing
outside that folder.

## For Users

```bash
# Clone and cd in
git clone <repo-url>
cd interesting-interview-questions

# Run tests for a single question
cd questions/easy/two-sum/go
go test ./...

# List available questions
go run ./cmd/report/

# Run every question's tests (from the repo root)
./bin/test

# Show question text and solution for specific ones
go run ./cmd/report/ -f two-sum,fizzbuzz -s

# Filter by difficulty
go run ./cmd/report/ -d easy,medium
```

1. Pick a question (`go run ./cmd/report/`).
2. Open `questions/<difficulty>/<short-name>/README.md` — the problem
   statement.
3. Implement your solution in the stub inside
   `questions/<difficulty>/<short-name>/go/` (a runnable Go module).
4. Run the tests (`cd questions/<difficulty>/<short-name>/go && go test ./...`).
5. Green? Check the official solution in `questions.yaml` (under `solution`).
6. Red? Debug and try again.

## For Maintainers

`questions.yaml` is the single source of truth for question metadata
and Go function signatures. From it, `cmd/generate` produces each
problem's `README.md`, `go.mod`, and stub. Stub code and `README.md`
are generated artifacts — never edit them by hand, and never put a
solution in a stub (solutions live only in `questions.yaml`).

```bash
# Regenerate stubs/READMEs and verify everything (builds pass, tests fail)
./bin/sync

# Overwrite implemented stubs during regeneration
FORCE=1 ./bin/sync

# Run every question's test suite
./bin/test
```

To add a question:

1. Add an entry to `questions.yaml`
   (see existing entries for the `signature` / `kind` format).
2. Add a `solution` field to the entry in `questions.yaml`.
3. Write a test file at
   `questions/<difficulty>/<short-name>/go/<short-name>_test.go`
   (`package main`, same folder as the stub — the real tests that verify
   a solved implementation).
4. Run `./bin/sync` to generate the `README.md`, `go.mod`, and stub.
5. Commit `questions.yaml` and the new directory.

## Difficulty Scale

Questions progress from straightforward loops to complex design problems.
Each directory in `questions/` is one of four levels:

| Level | Example | What to expect |
|-------|---------|----------------|
| Easy  | Fizz Buzz (`questions/easy/fizzbuzz/`) | One loop, conditionals, modulo — no data structures |
| Medium | Product of Array Except Self (`questions/medium/product-array-except-self/`) | O(n) time, no division — needs a different angle |
| Hard  | LRU Cache (`questions/hard/lru-cache/`) | Design a data structure with O(1) get and put with eviction |
| Expert | Count of Smaller Numbers After Self (`questions/expert/count-smaller-numbers-after-self/`) | Fenwick tree or merge sort; complex O(n log n) |
