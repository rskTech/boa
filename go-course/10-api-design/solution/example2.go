package main
import "testing"
func Add(a,b int) int { return a+b }
func TestAdd(t *testing.T){ tests:=[]struct{a,b,exp int}{ {1,2,3},{2,3,5} }
for _,tc:=range tests{ if got:=Add(tc.a,tc.b); got!=tc.exp{ t.Fatalf("expected %d got %d",tc.exp,got) } } }
