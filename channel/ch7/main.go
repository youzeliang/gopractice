package main

import (
	"fmt"
)

func main() {

	ch := prodece(1, 2, 3, 4)
	otu := square(ch)
	for n := range otu {
		fmt.Println(n)
	}

}

func prodece(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()

	return out
}

func square(inCh <-chan int) <-chan int {

	out := make(chan int)

	go func() {
		defer close(out)
		for n := range inCh {
			out <- n * n
		}
	}()

	return out
}
