// commands to run in terminal:
// go mod init example.com/myapp
// go get github.com/google/uuid@v1.3.0
package main
import (
    "fmt"
    "github.com/google/uuid"
)
func main(){ fmt.Println(uuid.New()) }
