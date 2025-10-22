package main
import "fmt"
type Book struct{ Title string; Pages int }
func main(){ m:=map[string]Book{"go":{"The Go Way",300}}; b:=m["go"]; fmt.Println(b.Title) }
