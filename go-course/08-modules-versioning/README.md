# 📦 Go Modules and Versioning

Go Modules is the **dependency management system** in Go, introduced in Go 1.11. It allows you to **manage project dependencies**, **version your code**, and build reproducible projects.

---

## 1️⃣ What is a Go Module?

- A **module** is a collection of Go packages stored in a **repository** with a `go.mod` file at its root.
- It defines the module’s **path** and **dependencies**.
- Eliminates the need for `GOPATH` for managing projects.

**Benefits:**

- Handles dependencies automatically.
- Supports semantic versioning (`v1.0.0`).
- Makes project reproducible across machines.

---

## 2️⃣ Initializing a Module

### Step 1: Create a project folder

```bash
mkdir go-project
cd go-project
```
Step 2: Initialize Go Module
```
go mod init github.com/username/go-project
```

Explanation:

go mod init creates a go.mod file.

The argument is the module path (usually your repository path).

Expected go.mod:
```
module github.com/username/go-project

go 1.24
```
## 3️⃣ Adding Dependencies
Step 1: Use a package in your code
```
package main

import (
    "fmt"
    "github.com/fatih/color"
)

func main() {
    color.Cyan("Hello, Go Modules!")
}
```
Step 2: Download dependencies
```
go mod tidy

```
Explanation:

go mod tidy downloads and adds missing modules to go.mod.

It also removes unused dependencies.

Updated go.mod:
```
module github.com/username/go-project

go 1.24

require github.com/fatih/color v1.13.0
```
## 4️⃣ Updating Dependencies
```
go get github.com/fatih/color@v1.14.0
```

Explanation:

go get updates the dependency to a specific version.

Supports semantic versioning.

Check versions:
```
go list -m all
```
## 5️⃣ Removing Dependencies

Remove the import from code.

Run:
```
go mod tidy
```

Unused dependencies are removed from go.mod and go.sum.

## 6️⃣ Go Module Best Practices

Always use modules for new projects.

Use semantic versioning for your own packages.

Run go mod tidy regularly to keep dependencies clean.

Commit go.mod and go.sum to version control.

Avoid manually editing go.sum.

## Notes
| Command                 | Purpose                                    |
| ----------------------- | ------------------------------------------ |
| `go mod init`           | Initialize a new module                    |
| `go mod tidy`           | Add missing and remove unused dependencies |
| `go get <module>@<ver>` | Add or update dependency                   |
| `go list -m all`        | List all module dependencies               |
| `go mod vendor`         | Copy dependencies into vendor folder       |

## Exercises
- Create a new module and import github.com/fatih/color to print colored text.

- Update the imported package to the latest version using go get.

- Remove the dependency and ensure go.mod is updated correctly.

- Initialize a module in an existing project and verify reproducible builds using go mod tidy
