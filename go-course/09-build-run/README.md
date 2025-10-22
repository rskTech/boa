# 🏗️ Building and Running Go Programs

In Go, building and running programs is simple thanks to its **toolchain**. Understanding how to compile, run, and manage Go binaries is essential for development and deployment.

---

## 1️⃣ Running Go Programs with `go run`

- The simplest way to execute a Go program is using `go run`.

### Example

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go World!")
}
```
Run the program:
```
go run main.go
```

Expected Output:
```
Hello, Go World!
```

Explanation:

go run compiles and immediately runs the program.

Useful during development and testing.

## 2️⃣ Building Executables with go build

Use go build to create an executable binary.

Example
```
go build main.go
```

Creates an executable file named main (or main.exe on Windows).

Run the binary:
```
./main    # Linux/macOS
main.exe  # Windows
```

Output:
```
Hello, Go World!
```

Explanation:

go build compiles your program and creates a standalone binary.

No source code is needed to run the binary.

## 3️⃣ Build Flags and Options

```
-o <filename>: Specify output file name.

go build -o hello main.go
./hello
```

-v: Verbose mode, shows packages being compiled.
```
go build -v main.go
```

-a: Force rebuilding of all packages.
```
go build -a main.go
```
## 4️⃣ Running Programs in Module Context

With Go modules, ensure you run commands inside the module directory (where go.mod is present):
```
go run ./cmd/app
go build ./cmd/app
```

This ensures correct dependency resolution.

## 5️⃣ Cross-Compilation

Go supports building binaries for different OS and architectures.
```
GOOS=linux GOARCH=amd64 go build -o myapp-linux main.go
GOOS=windows GOARCH=amd64 go build -o myapp.exe main.go
```

Explanation:

GOOS: Target operating system (linux, windows, darwin)

GOARCH: Target architecture (amd64, arm64)

Creates a binary compatible with the specified platform.

## 6️⃣ Running Tests While Building

Go provides built-in testing:
```
go test ./...
```

Run all tests in the module recursively.

Use -v for verbose output:
```
go test -v ./...
```
