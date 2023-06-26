package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(A())

	time.Sleep(4 * time.Second)
}

func A() int {

	go func() {
		//time.Sleep(time.Second*3)
		fmt.Println(43241)
	}()

	return 14
}
