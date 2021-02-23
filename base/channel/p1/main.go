package main

import (
	"fmt"
)

func goroutineA(a <-chan int) {
	val := <-a
	fmt.Println("goroutine A received data: ", val)
	return
}

func goroutineB(b <-chan int) {
	val := <-b
	fmt.Println("goroutine B received data: ", val)
	return
}

func main() {
	//ch := make(chan int)
	//go goroutineA(ch)
	//go goroutineB(ch)
	//ch <- 3
	//time.Sleep(time.Second)
	//
	//ch1 := make(chan struct{})
	//fmt.Println(ch1)

	fmt.Printf("%d\n")
	return
	fmt.Println("Done")
}
