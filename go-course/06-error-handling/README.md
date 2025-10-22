# ⚠️ Error Handling in Go

Go handles errors **explicitly** rather than using exceptions like other languages. Understanding error handling is crucial for writing **robust and reliable programs**.

---

## 1️⃣ Understanding Errors

In Go, **errors are values** returned by functions. They implement the built-in `error` interface:

```go
type error interface {
    Error() string
}
```
- If a function can fail, it usually returns (result, error).

- Callers check the error value and handle it accordingly.

## 2️⃣ Basic Error Handling
Example 1: Handling a Division Function
```
package main

import (
    "errors"
    "fmt"
)

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

func main() {
    result, err := divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result)
    }
}

```
Expected Output:
```
Error: division by zero
```

Explanation:

Function returns both result and error.

Caller checks err before using the result.

## 3️⃣ Using fmt.Errorf for Custom Messages

fmt.Errorf allows formatting error messages:
```
package main

import "fmt"

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide %v by zero", a)
    }
    return a / b, nil
}

func main() {
    result, err := divide(15, 0)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("Result:", result)
    }
}
```

Output:
```
cannot divide 15 by zero
```
## 4️⃣ Panic and Recover
### Panic

Used for unexpected, unrecoverable errors.

Stops normal execution.
```
package main
import "fmt"

func main() {
    panic("Something went wrong!")
    fmt.Println("This will not be printed")
}

```
Output:
```
panic: Something went wrong!
```
### Recover

Allows program to recover from panic and continue.
```
package main
import "fmt"

func safeDivide(a, b float64) {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()

    if b == 0 {
        panic("division by zero")
    }
    fmt.Println("Result:", a/b)
}

func main() {
    safeDivide(10, 0)
    fmt.Println("Program continues...")
}

```
Output:
```
Recovered from panic: division by zero
Program continues...
```
## 5️⃣ Best Practices

Return errors explicitly; do not ignore them.

Use errors.New or fmt.Errorf for custom error messages.

Use panic only for critical, unrecoverable errors.

Use recover in deferred functions to gracefully handle panics.

Always check for errors from I/O, network, or database operations.

## Notes
| Concept      | Purpose                                 |
| ------------ | --------------------------------------- |
| `error`      | Standard error type in Go               |
| `errors.New` | Create a simple error                   |
| `fmt.Errorf` | Create formatted/custom error messages  |
| `panic`      | Stop execution due to unexpected error  |
| `recover`    | Recover from panic in deferred function |

## Exercises 
- Create a function that opens a file and returns an error if the file does not exist.

- Modify the division function to log errors to the console.

- Write a function that intentionally panics and use recover to continue program execution.

- Experiment with multiple nested function calls where panic occurs and recover handles it.
