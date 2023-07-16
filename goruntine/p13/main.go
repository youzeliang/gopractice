package main

import (
	"fmt"
	"time"
)

var ch = make(chan string, 1)

func producer() {
	for i := 0; i < 10; i++ {
		go func(pi int) {
			n := 0
			for {
				n += 1
				ch <- fmt.Sprintf("%d_%d", pi, n)
				fmt.Printf("%s producer write p:%d, n:%d\n", time.Now().String(), pi, n)
				time.Sleep(1 * time.Second)
			}
		}(i)
	}
}

func consumer() {
	for {
		gotS := <-ch
		fmt.Printf("%s consumer read n:%s\n", time.Now().String(), gotS)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go consumer()
	go producer()
	select {}
}
