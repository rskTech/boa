package main
import "fmt"
func main(){
    m := map[string]int{"a":1, "b":2}
    m["c"] = 3
    for k,v := range m { fmt.Println(k,v) }
}
