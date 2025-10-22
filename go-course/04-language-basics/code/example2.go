package main
import "fmt"
func div(a,b float64)(float64,error){ if b==0{return 0, fmt.Errorf("divide by zero")} return a/b,nil }
func main(){ if r,err:=div(10,2); err==nil{ fmt.Println(r) } }
