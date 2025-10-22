# Lesson 03 — Setting Up Go Development Environment (VS Code)

Steps (Ubuntu):
1. Install Go: `sudo apt install -y golang-go`
2. Verify: `go version`
3. Install VS Code and the Go extension (`golang.go`)
4. Open the lesson folder in VS Code and accept tool installations (gopls, goimports, delve...)

Example: create a module and run a program
```bash
mkdir -p lessons/03/code && cd lessons/03/code
go mod init example.com/lesson3
go run main.go
```
