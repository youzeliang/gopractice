package main

import (
	"fmt"
	"time"
)

//func main() {
//	timeOut()
//}

func timeOut() {
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(time.Second) // sleep one second
		timeout <- true
	}()
	ch := make(chan int)
	select {
	case <-ch:
	case <-timeout:
		fmt.Println("timeout!")
	}
}
func main() {

	ch := make(chan int, 1)
	ch <- 1
	select {
	case ch <- 2:
	default:
		fmt.Println("channel is full !")
	}

}
