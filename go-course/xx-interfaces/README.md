# 🔗 Interfaces in Go

In Go, **interfaces** provide a way to define **behavior** without specifying **implementation**. They are central to Go’s approach to polymorphism and allow writing flexible and extensible programs.

---


## 1️⃣ What is an Interface?

An interface is a **set of method signatures**. Any type that implements those methods **satisfies the interface**. Go does not require explicit declaration of intent to implement an interface — it is **implicit**.

### Syntax
```go
type InterfaceName interface {
    Method1(paramType) returnType
    Method2(paramType) returnType
}
```
## 2️⃣ Example 1: Simple Interface
```
package main
import "fmt"

// Define an interface
type Shape interface {
    Area() float64
}

// Implement interface for Rectangle
type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Implement interface for Circle
type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

func main() {
    var s Shape

    r := Rectangle{Width: 5, Height: 4}
    c := Circle{Radius: 3}

    s = r
    fmt.Println("Rectangle Area:", s.Area())

    s = c
    fmt.Println("Circle Area:", s.Area())
}
```

Expected Output:
```
Rectangle Area: 20
Circle Area: 28.26
```

Explanation:

Both Rectangle and Circle implement the Shape interface.

Interface allows polymorphic behavior — the same variable s can hold any type that implements Shape.

## 3️⃣ Example 2: Interface with Multiple Methods
```
package main
import "fmt"

type Vehicle interface {
    Start()
    Stop()
}

type Car struct {
    Brand string
}

func (c Car) Start() {
    fmt.Println(c.Brand, "started")
}

func (c Car) Stop() {
    fmt.Println(c.Brand, "stopped")
}

func main() {
    var v Vehicle = Car{Brand: "Toyota"}
    v.Start()
    v.Stop()
}

```
Expected Output:
```
Toyota started
Toyota stopped
```

Explanation:

Interfaces can have multiple methods.

Any type implementing all methods satisfies the interface.

## 4️⃣ Empty Interface (interface{})

An empty interface can hold values of any type.

Useful for generic programming.
```
package main
import "fmt"

func printValue(value interface{}) {
    fmt.Println(value)
}

func main() {
    printValue(42)
    printValue("Hello")
    printValue(3.14)
}

```
Expected Output:
```
42
Hello
3.14
```
## 5️⃣ Key Points

Go uses implicit implementation — no implements keyword.

Interfaces are satisfied automatically if the type has required methods.

Interfaces enable polymorphism, flexibility, and decoupling of code.

Empty interfaces (interface{}) can store any type of value.

Use type assertions or type switches to retrieve the actual type from an interface value.

## Notes
| Concept         | Description                                        |
| --------------- | -------------------------------------------------- |
| Interface       | Set of method signatures                           |
| Satisfies       | Type implements all methods in the interface       |
| Polymorphism    | Single interface variable can hold multiple types  |
| Empty Interface | `interface{}` can store any value                  |
| Type Assertion  | Retrieve the concrete type from an interface value |

## Exercises 
- Define an interface Animal with method Speak(). Implement it for Dog and Cat.

- Write a function that takes interface{} as input and prints the type and value.

- Create an interface Notifier with methods SendEmail and SendSMS. Implement it for two different struct types.

- Use type assertion to identify the actual type stored in an interface variable.
