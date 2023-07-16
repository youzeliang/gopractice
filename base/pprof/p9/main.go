package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- 42
	}()

	select {
	case data := <-ch:
		fmt.Println("读取到数据:", data)
	case <-time.After(3 * time.Second):
		fmt.Println("超时，未能读取到数据")
	}

	time.Sleep(time.Second * 5)
}
