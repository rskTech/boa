# 🛠️ Setting Up the Go Development Environment

This module will guide you through **installing Go, setting up your workspace, and configuring your development environment** so you can start building Go programs efficiently.

## 🧩 Step 1: Install Go

### Linux / Ubuntu

```bash
sudo apt update
sudo apt install -y golang
```
Verify installation:
```
go version
```
Expected Output Example:
```
go version go1.21 linux/amd64
```
macOS

Using Homebrew:
```
brew install go
```
Verify:
```
go version
```
Windows

Download the Go installer from https://go.dev/dl/.
Run the installer and follow instructions.
Verify in PowerShell or CMD:
```
go version
```
## 🧩 Step 2: Set Up Environment Variables
Linux / macOS

Add to ~/.bashrc or ~/.zshrc:
```
export GOPATH=$HOME/go
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
```
Apply changes:
```
source ~/.bashrc   # or source ~/.zshrc
```
Verify:
```
echo $GOPATH
```
Expected Output Example:
```
/home/username/go
```
Windows

Set GOPATH as a system environment variable (e.g., C:\Users\<YourUser>\go).

Add %GOPATH%\bin and Go installation path (C:\Go\bin) to the system PATH.

Restart terminal and verify:
```
go env GOPATH
```
## 🧩 Step 3: Create a Go Workspace
Create directories for Go workspace:
```
mkdir -p $GOPATH/src
mkdir -p $GOPATH/bin
mkdir -p $GOPATH/pkg
```
Workspace Structure:
```
go-workspace/
├── bin/       # Compiled binaries
├── pkg/       # Packages
└── src/       # Source code
```
## 🧩 Step 4: Configure VS Code for Go
Install Visual Studio Code.

Install Go extension by Microsoft.

Open any Go file → VS Code will prompt to install Go tools (gopls, goimports, etc.) → Install them.

Set the workspace folder to your Go project directory.

## 🧩 Step 5: Verify the Setup
Create a file main.go in your workspace:
```
package main

import "fmt"

func main() {
    fmt.Println("Go environment is ready!")
}
```
Run the program:
```
go run main.go
```
Expected Output:
```
Go environment is ready!
```
Compile to Binary (Optional)
```
go build main.go
./main   # Run the compiled binary
```
