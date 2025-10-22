package main
import (
    "fmt"
    "time"
)
func main() {
    fmt.Println("Hello, Go!")
    fmt.Println("Time:", time.Now().Format(time.RFC1123))
}
