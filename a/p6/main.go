package main

import (
	"fmt"
	"time"
)

// 写未初始化的chan
func main() {

	c := make(chan int)
	go func() {
		time.Sleep(time.Second * 1)
		c <- 10
		close(c)
	}()

	for {
		select {
		case x, ok := <-c:
			fmt.Println(x, ok)
			time.Sleep(time.Second * 3)
		default:
			fmt.Println("no")
		}
	}
}
