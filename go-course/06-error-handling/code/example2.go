package main
import (
    "errors"
    "fmt"
)
var ErrNotFound = errors.New("not found")
func find(x int) error { return fmt.Errorf("find failed: %w", ErrNotFound) }
func main(){ err:=find(1); if errors.Is(err, ErrNotFound){ fmt.Println("not found") } }
