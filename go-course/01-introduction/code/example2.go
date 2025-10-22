package main
import (
    "fmt"
    "os"
)
func main() {
    if len(os.Args) > 1 {
        fmt.Println("Arg:", os.Args[1])
    } else {
        fmt.Println("No arg provided")
    }
}
