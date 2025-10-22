# ⚡ Concurrency Basics in Go: Goroutines and Channels

Go is designed with **concurrency in mind**. Goroutines and channels are the core constructs that make Go programs **efficient, scalable, and concurrent**.

---

## 1️⃣ Goroutines

### What is a Goroutine?

- A **goroutine** is a lightweight thread managed by the Go runtime.
- Allows functions to run **concurrently** with other functions.
- Much cheaper than traditional threads.

### Syntax

Example 1: Basic Goroutine
```
package main
import (
    "fmt"
    "time"
)

func sayHello() {
    fmt.Println("Hello from goroutine")
}

func main() {
    go sayHello()   // run as goroutine
    time.Sleep(time.Second) // wait for goroutine to finish
    fmt.Println("Main function ends")
}

```
Expected Output:
```
Hello from goroutine
Main function ends
```

Explanation:

go sayHello() runs concurrently.

time.Sleep is used to give time for the goroutine to finish (not ideal for production).

Example 2: Multiple Goroutines
```
package main
import (
    "fmt"
    "time"
)

func printNumbers(prefix string) {
    for i := 1; i <= 5; i++ {
        fmt.Println(prefix, i)
        time.Sleep(100 * time.Millisecond)
    }
}

func main() {
    go printNumbers("Goroutine1")
    go printNumbers("Goroutine2")
    time.Sleep(time.Second)
    fmt.Println("Main ends")
}

```
Expected Output (order may vary due to concurrency):
```
Goroutine1 1
Goroutine2 1
Goroutine1 2
Goroutine2 2
...
Main ends
```
## 2️⃣ Channels
What is a Channel?

Channels are pipes that allow goroutines to communicate and synchronize.

A channel can send and receive values of a specific type.

Syntax
```
ch := make(chan int) // unbuffered channel
```
Example 1: Sending and Receiving
```
package main
import "fmt"

func greet(ch chan string) {
    ch <- "Hello from goroutine" // send message to channel
}

func main() {
    message := make(chan string)
    go greet(message)
    msg := <-message // receive message from channel
    fmt.Println(msg)
}

```
Expected Output:
```
Hello from goroutine
```
Example 2: Buffered Channels
```
package main
import "fmt"

func main() {
    ch := make(chan int, 3) // buffered channel of size 3

    ch <- 1
    ch <- 2
    ch <- 3

    fmt.Println(<-ch)
    fmt.Println(<-ch)
    fmt.Println(<-ch)
}

```
Expected Output:
```
1
2
3
```

Explanation:

Buffered channels allow sending multiple values without immediate receiver.

Unbuffered channels require simultaneous send and receive.

Example 3: Using Channels for Synchronization
```
package main
import "fmt"

func worker(done chan bool) {
    fmt.Println("Working...")
    done <- true // signal completion
}

func main() {
    done := make(chan bool)
    go worker(done)
    <-done // wait for worker to finish
    fmt.Println("Worker finished")
}

```
Output:
```
Working...
Worker finished
```

Explanation:

Channels can be used to coordinate goroutines and ensure tasks complete before proceeding.

## 3️⃣ Key Points

Goroutines are lightweight threads; use go keyword.

Channels are the safest way to communicate between goroutines.

Always close channels if no more data will be sent.

Use buffered channels for better performance with multiple sends.

Avoid sharing data using variables; prefer channels for communication.

## Notes
| Concept            | Description                                      |
| ------------------ | ------------------------------------------------ |
| Goroutine          | Lightweight concurrent function                  |
| Channel            | Communication pipe between goroutines            |
| Buffered Channel   | Channel with capacity, doesn't block immediately |
| Unbuffered Channel | Send blocks until received                       |
| Synchronization    | Use channels to signal completion                |

## Exercises 
- Launch 3 goroutines to print numbers 1-5. Ensure main waits for all to finish.

- Create a buffered channel of size 5 and send numbers to it. Receive and print all numbers.

- Write a function that calculates squares of numbers concurrently using goroutines and channels.

- Use multiple goroutines to fetch data and synchronize results using channels.
