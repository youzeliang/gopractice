package main

import (
	"fmt"
	"sync"
)

func main() {
	a := make([]int, 0)
	buffer := make(chan int)
	go func() {
		for v := range buffer {
			a = append(a, v)
		}
	}()

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			buffer <- i
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println(len(a))
}
