package main

import (
	"fmt"
	"reflect"
	"time"
)

var ch = make(chan int)

func main() {
	for i := 1; i <= 75; i++ {
		go print(i)
		ch <- 0
	}
	time.Sleep(5 * time.Second) //wait all goroutines
}

func print(i int) {
	<-ch
	fmt.Println(i)

}
