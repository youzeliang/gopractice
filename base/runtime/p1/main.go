package main

import (
	"fmt"
	"runtime"
	"time"
)

func a() {
	for i := 1; i < 10; i++ {
		time.Sleep(time.Second * 1)
		fmt.Println("A:", i)
	}
}
func b() {
	for i := 1; i < 10; i++ {
		time.Sleep(time.Second * 1)
		fmt.Println("B:", i)
	}
}
func main() {
	runtime.GOMAXPROCS(1)
	go a()
	go b()
	time.Sleep(130 * time.Second)
}
