# 📝 Go Language Basics

In this module, we will cover the **fundamental building blocks of the Go programming language**, including variables, data types, control structures, functions, and packages. Understanding these concepts is crucial for writing clean and efficient Go programs.

---

## 1️⃣ Variables and Data Types

### What are Variables?

Variables are **named storage locations** in memory that hold values. In Go, every variable has a **type** which defines what kind of value it can store.

### Variable Declaration

1. **Explicit Declaration**
```go
var name string = "Rajendra"
var age int = 25
```
2. **Type Inference**
```
name := "Rajendra"
age := 25
```

3. **Multiple Variables**
```
var x, y, z int = 1, 2, 3
a, b := 10, "Hello"
```
4. **Common Data Types**

| Type    | Description            | Example      |
| ------- | ---------------------- | ------------ |
| int     | Integer                | 10           |
| float64 | Floating-point number  | 3.14         |
| string  | Sequence of characters | "GoLang"     |
| bool    | Boolean                | true / false |
| byte    | Alias for uint8        | 255          |
| rune    | Unicode code point     | 'a'          |


Example 1: Using Variables
```package main
import "fmt"

func main() {
    var name string = "GoLang"
    age := 15
    fmt.Println("Name:", name, "Age:", age)
}
```

Expected Output:
```
Name: GoLang Age: 15
```
Example 2: Multiple Variables
```package main
import "fmt"

func main() {
    var x, y int = 5, 10
    sum := x + y
    fmt.Println("Sum:", sum)
}
```
## 2️⃣ Control Structures (If-Else, Loops)
### If-Else Statement

Used to execute code conditionally.
```
package main
import "fmt"

func main() {
    age := 18
    if age >= 18 {
        fmt.Println("You are an adult")
    } else {
        fmt.Println("You are a minor")
    }
}
```

Expected Output:
```
You are an adult
```
### For Loop

Go has only the for loop (no while or do-while).

Simple for loop
```
for i := 1; i <= 5; i++ {
    fmt.Println(i)
}
```

For loop as while
```
i := 1
for i <= 5 {
    fmt.Println(i)
    i++
}
```

Infinite loop
```
for {
    fmt.Println("Infinite loop")
    break
}
```
Example: Sum of Numbers
```
package main
import "fmt"

func main() {
    sum := 0
    for i := 1; i <= 5; i++ {
        sum += i
    }
    fmt.Println("Sum of 1 to 5:", sum)
}
```

Expected Output:
```
Sum of 1 to 5: 15
```
## 3️⃣ Functions and Packages
### Functions

Functions allow us to reuse code and structure programs.

Syntax:
```
func functionName(parameters) returnType {
    // code
}
```
Example 1: Simple Function
```
package main
import "fmt"

func greet(name string) {
    fmt.Println("Hello,", name)
}

func main() {
    greet("Rajendra")
}
```

Expected Output:
```
Hello, Rajendra
```
Example 2: Function with Return
```
package main
import "fmt"

func add(a int, b int) int {
    return a + b
}

func main() {
    sum := add(10, 15)
    fmt.Println("Sum:", sum)
}
```

Expected Output:
```
Sum: 25
```
### Packages

Go programs are organized in packages.

main package is used for the entry point of the program.

Other packages can be imported using import.

Example: Using Math Package
```
package main
import (
    "fmt"
    "math"
)

func main() {
    fmt.Println("Square root of 16 is", math.Sqrt(16))
}

```
Expected Output:
```
Square root of 16 is 4
```
## ⚡ Quick Notes

Variables declared but not used will cause a compile-time error.

Go has no ternary operator; use if-else.

Only the first letter capitalized functions or variables are exported from packages.

## 🧩 Exercises

- Declare variables of different types and print their values.

- Write a program to find the largest of 3 numbers using if-else.

- Create a function that calculates the factorial of a number.

- Use a loop to print even numbers from 1 to 20.

- Import the strings package and convert a string to uppercase.
