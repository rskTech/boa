package main
import "fmt"
type Person struct {
    Name string
    Age int
}
func (p Person) Greet(){ fmt.Println("Hi, I am", p.Name) }
func main(){ p:=Person{"Raj",34}; p.Greet() }
