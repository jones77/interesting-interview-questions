# Interesting Interview Questions

Clone this repo, implement your own solutions, and make the tests pass.
Reference solutions are in `questions.yaml` for when you want to check
your work — solve it first, then peek.

## For Users

```bash
# Clone and cd in
git clone <repo-url>
cd interesting-interview-questions

# Run tests for all questions
go test ./...

# Run tests for a single question
go test ./questions/easy/two-sum/

# List available questions
go run ./cmd/report/

# Show question text and solution for specific ones
go run ./cmd/report/ -f two-sum,fizzbuzz -s

# Filter by difficulty
go run ./cmd/report/ -d easy,medium
```

1. Pick a question (`go run ./cmd/report/`).
2. Open the stub — the problem statement is in the doc comment.
3. Implement your solution in the function body.
4. Run the tests (`go test ./questions/<difficulty>/<short-name>/`).
5. Green? Check the official solution in `questions.yaml` (under `solution`).
6. Red? Debug and try again.

## For Maintainers

`questions.yaml` is the single source of truth for question metadata
and Go function signatures. Stub `.go` files are generated from it.

```bash
# Regenerate stubs and verify everything builds
./bin/sync

# Overwrite implemented stubs during regeneration
FORCE=1 ./bin/sync
```

To add a question:

1. Add an entry to `questions.yaml`
   (see existing entries for the `signature` / `kind` format).
2. Add a `solution` field to the entry in `questions.yaml`.
3. Write a test file at
   `questions/<difficulty>/<short-name>/<short-name>_test.go`.
4. Run `./bin/sync` to regenerate the stub.
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
