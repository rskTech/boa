            package main
            import (
                "errors"
	"fmt"
)
            func mayfail(x int) (int, error) {
                if x<0 { return 0, errors.New("negative not allowed") }
                return x*2, nil
            }
            func main() {
                if v, err := mayfail(-1); err != nil { fmt.Println("error:", err) } else { fmt.Println(v) }
            }
