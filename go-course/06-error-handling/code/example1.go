package main
import (
    "errors"
    "fmt"
)
func mayFail(x int)(int,error){ if x<0{ return 0, errors.New("negative") }; return x*2,nil }
func main(){ if v,err:=mayFail(-5); err!=nil{ fmt.Println("err:",err) } else { fmt.Println(v) } }
