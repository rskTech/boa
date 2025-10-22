package main
import "fmt"
func main(){ s:=[]int{1,2}; s=append(s,3,4); t:=make([]int,len(s)); copy(t,s); fmt.Println(s,t) }
