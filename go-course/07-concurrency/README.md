# Lesson 07 — Concurrency Basics (Goroutines and Channels)

## Concurrency
- Goroutines are cheap; channels are communication primitives; use `select` to multiplex.
- Use `-race` detector to find data races during tests.


## Examples
- `code/example1.go` — Worker pool pattern
- `code/example2.go` — Select with timeout

See `exercises.md` for tasks and hints.
