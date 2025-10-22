# 🧬 Inheritance in Go

Go does not support classical object-oriented inheritance (like Java or C++). Instead, it uses **composition** and **interfaces** to achieve similar functionality. Understanding inheritance in Go is essential for designing flexible and maintainable programs.

---

## 1️⃣ Composition in Go (Struct Embedding)

- Instead of inheriting from a parent class, Go allows **embedding one struct inside another**.  
- The embedded struct's fields and methods are accessible directly from the outer struct.

### Example 1: Basic Struct Embedding

```go
package main

import "fmt"

// Base struct
type Animal struct {
    Name string
}

func (a Animal) Speak() {
    fmt.Println(a.Name, "makes a sound")
}

// Derived struct via embedding
type Dog struct {
    Animal // Embedding Animal struct
    Breed  string
}

func main() {
    d := Dog{
        Animal: Animal{Name: "Buddy"},
        Breed:  "Golden Retriever",
    }
    
    // Access embedded struct method
    d.Speak()
    
    // Access embedded struct field
    fmt.Println("Dog Breed:", d.Breed)
}
```
Expected Output:
```
Buddy makes a sound
Dog Breed: Golden Retriever
```

Explanation:

Dog struct embeds Animal struct.

Dog can use Animal’s methods and fields directly.

This achieves a form of inheritance.

Example 2: Method Overriding

Go does not support traditional method overriding, but you can define a method with the same name in the outer struct to override behavior.
```
func (d Dog) Speak() {
    fmt.Println(d.Name, "barks")
}

func main() {
    d := Dog{
        Animal: Animal{Name: "Buddy"},
        Breed:  "Golden Retriever",
    }
    d.Speak() // Calls Dog's Speak, not Animal's
}
```

Expected Output:
```
Buddy barks
```

Explanation:

Defining Speak() in Dog “overrides” Animal’s Speak() for Dog.

## 2️⃣ Interface-Based Polymorphism

Interfaces allow polymorphic behavior, similar to inheritance in other languages.

Any type that implements the methods of an interface satisfies it automatically.

Example 3: Interface Implementation
```
package main

import "fmt"

// Define an interface
type Speaker interface {
    Speak()
}

// Base struct
type Cat struct {
    Name string
}

func (c Cat) Speak() {
    fmt.Println(c.Name, "meows")
}

func main() {
    var s Speaker
    c := Cat{Name: "Whiskers"}
    
    s = c
    s.Speak() // Calls Cat's Speak method
}

```
Expected Output:
```
Whiskers meows

```
Explanation:

Speaker interface defines a Speak() method.

Cat struct implements Speak(), so it satisfies Speaker.

This allows dynamic behavior similar to class inheritance in other languages.

## ⚡ Key Notes

Go favors composition over inheritance.

Struct embedding provides reusable fields and methods.

Interfaces enable polymorphism without tight coupling.

Avoid deep inheritance hierarchies; use small, composable structs and interfaces.

## 🧩 Exercises

- Create a Vehicle struct with a method Move(). Embed it into a Car struct and call Move().

- Override Move() in Car to provide specific behavior.

- Define a Mover interface with method Move(). Ensure both Vehicle and Car satisfy it.

- Demonstrate polymorphism by writing a function that accepts any Mover and calls Move().
