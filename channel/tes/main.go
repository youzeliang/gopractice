package main

import "fmt"

func main() {
	a()
}

func a() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for {
			m, ok := <-ch1
			if !ok {
				break
			}

			ch2 <- m
		}
		close(ch2)
	}()

	for i := range ch2 {
		fmt.Println(i)
	}
}
