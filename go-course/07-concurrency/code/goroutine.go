            package main
            import (
	"fmt"
	"time"
)

            func worker(id int, ch chan int) {
	for i:=0;i<3;i++ { ch <- id*10 + i }
}

            func main(){ ch := make(chan int)
	for i:=0;i<2;i++ { go worker(i,ch) }
	for i:=0;i<6;i++ { fmt.Println(<-ch) }
	close(ch)
	time.Sleep(100*time.Millisecond)
}
