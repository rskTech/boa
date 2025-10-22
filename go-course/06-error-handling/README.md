# Lesson 06 — Error Handling in Go

## Error Handling
- Prefer returning errors; use `errors` package and wrapping with `%w`.
- Use `errors.Is`/`errors.As` to inspect wrapped errors.
- Avoid panic for recoverable errors.


## Examples
- `code/example1.go` — Basic error return
- `code/example2.go` — Wrapping and inspecting errors

See `exercises.md` for tasks and hints.
