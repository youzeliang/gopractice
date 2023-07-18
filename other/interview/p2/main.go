package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1000)
	go func() {
		for i := 0; i < 1; i++ {
			ch <- i
		}
	}()
	go func() {
		for {
			a, ok := <-ch
			if !ok {
				fmt.Println("close")
				return
			}
			fmt.Println("a: ", a)
		}
	}()
	time.Sleep(time.Second * 1)
	close(ch)
	fmt.Println("ok")
}

type S struct {
}

func f(x interface{}) {
}

func g(x *interface{}) {
}
