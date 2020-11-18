package main

import (
	"fmt"
	"time"
)

func main() {

	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)

	go func() {
		//time.Sleep(2 * time.Second)
		int_chan <- 1
	}()

	go func() {
		time.Sleep(2 * time.Second)
		string_chan <- "hello"
	}()

	select {
	case value := <-int_chan:
		fmt.Println("int:", value)
	case value := <-string_chan:
		fmt.Println("string:", value)
	}
	fmt.Println("main结束")

}
