# AGENTS.md

## Structure

Each interview question lives in a single `<short-name>.md` file at repo root — e.g. `two-sum.md`, `fizzbuzz.md`.

## Convention

- **Never include the solution** in a question's `.md` file. The file is the question.
- A **rot13 hint** is permissible if hints are traditionally given for that problem.
- Solutions may eventually live in `solutions/` (not yet active; `.gitignore` already ignored).

## Adding a Question

1. Create `{short-name}.md` with the problem statement only.
2. If a hint is appropriate, append a rot13-ciphered hint in a `> **Hint**` block.
3. Commit after each question for clean history.
