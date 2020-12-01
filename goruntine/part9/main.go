package main

import (
	"C"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	ch := make(chan int, 1)

	go Producer(2, ch)
	go Producer(3, ch)
	go Consumer(ch)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig,syscall.SIGINT,syscall.SIGTERM)

}

func Producer(faction int, out chan<- int) {

	for i := 0; ; i++ {
		out <- faction * i
	}
}

func Consumer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}
