# 📚 Data Structures in Go

In this module, we will cover the **essential data structures** in Go: **Arrays, Slices, Maps, and Structs**. Understanding these will help you organize and manage data efficiently in your programs.

---

## 1️⃣ Arrays

### What is an Array?

An array is a **fixed-size collection of elements of the same type**. The size must be known at compile time.

### Declaration

```go
var numbers [5]int             // array of 5 integers
names := [3]string{"Alice", "Bob", "Charlie"}
```
Accessing Elements
```
fmt.Println(names[0]) // Alice
numbers[2] = 10
```
Example 1: Sum of Array Elements
```
package main
import "fmt"

func main() {
    numbers := [5]int{1, 2, 3, 4, 5}
    sum := 0
    for i := 0; i < len(numbers); i++ {
        sum += numbers[i]
    }
    fmt.Println("Sum:", sum)
}

```
Output:
```
Sum: 15
```
Example 2: Iterating Over Array
```
for index, value := range numbers {
    fmt.Println(index, value)
}
```

Differences: 

Arrays are fixed-size, cannot grow dynamically, and copying an array creates a new copy.

## 2️⃣ Slices
### What is a Slice?

A slice is a dynamically sized, flexible view into an array. Most Go programs use slices instead of arrays.

Declaration
```
var nums []int                // nil slice
nums := []int{1, 2, 3, 4, 5} // initialized slice
```
Adding Elements
```
nums = append(nums, 6)
```
Example 1: Iterating Over Slice
```
package main
import "fmt"

func main() {
    nums := []int{10, 20, 30}
    nums = append(nums, 40)
    for _, num := range nums {
        fmt.Println(num)
    }
}
```

Output:
```
10
20
30
40
```
Example 2: Slicing a Slice
```
sub := nums[1:3]
fmt.Println(sub) // [20 30]
```

Differences:

Slices are dynamic in size.

Slices reference underlying arrays (modifying slice may modify array).

Preferred over arrays in most cases.

## 3️⃣ Maps
### What is a Map?

A map is a collection of key-value pairs. Keys are unique, and values are associated with keys.

Declaration
```
var studentGrades map[string]int
grades := map[string]int{"Alice": 90, "Bob": 85}
```
Example 1: Accessing and Updating Map
```
package main
import "fmt"

func main() {
    grades := map[string]int{"Alice": 90, "Bob": 85}
    fmt.Println("Alice's grade:", grades["Alice"])
    grades["Charlie"] = 95
    fmt.Println(grades)
}
```

Output:
```
Alice's grade: 90
map[Alice:90 Bob:85 Charlie:95]
```
Example 2: Iterating Over Map
```
for name, grade := range grades {
    fmt.Printf("%s scored %d\n", name, grade)
}
```

Differences:

Maps are unordered.

Keys must be comparable.

Very flexible for storing associative data.

## 4️⃣ Structs
### What is a Struct?

A struct is a custom data type that groups together fields of different types. It is used to model real-world entities.

Declaration
```
type Person struct {
    Name string
    Age  int
}
```
Example 1: Creating and Using Struct
```
package main
import "fmt"

type Person struct {
    Name string
    Age  int
}

func main() {
    p := Person{Name: "Rajendra", Age: 33}
    fmt.Println(p.Name, "is", p.Age, "years old")
}

```
Output:
```
Rajendra is 33 years old
```
Example 2: Struct with Function
```
func (p Person) Greet() {
    fmt.Printf("Hello, my name is %s\n", p.Name)
}

func main() {
    p := Person{"Ira", 5}
    p.Greet()
}

```
Output:
```
Hello, my name is Ira
```

Differences:

Structs can hold multiple fields of different types.

Used for modeling complex objects.

Supports methods (functions associated with structs).

## Comparison
| Feature     | Array      | Slice         | Map          | Struct            |
| ----------- | ---------- | ------------- | ------------ | ----------------- |
| Size        | Fixed      | Dynamic       | Dynamic      | N/A               |
| Type        | Same       | Same          | Key-Value    | Mixed fields      |
| Growth      | No         | Yes           | Yes          | N/A               |
| Order       | Ordered    | Ordered       | Unordered    | Ordered fields    |
| Example Use | Fixed data | Flexible list | Lookup table | Real-world entity |

## Exercises 
- Create an array of 5 integers and calculate their sum.

- Create a slice of strings, append 2 more elements, and print all elements.

- Create a map of student names to grades, update a grade, and print all entries.

- Define a struct Car with fields Brand, Model, and Year. Create an instance and print details.

- Add a method to Car struct that prints a summary of the car.
