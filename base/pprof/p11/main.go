package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 100)
	go func() {
		for {
			time.Sleep(time.Second * 1)
			ch <- "continue"
		}
	}()
	for {
		select {
		case c := <-ch:
			fmt.Println(c)

		case <-time.After(time.Minute * 2):
		}
	}
}
