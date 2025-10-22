            # Lesson 09 — Building and Running Go Programs

- `go run main.go` for quick runs
- `go build` to produce a binary
- Cross-compile with `GOOS` and `GOARCH`

Example:
```
GOOS=linux GOARCH=amd64 go build -o myapp
```
