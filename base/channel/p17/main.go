package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int, 2000)

	var wg sync.WaitGroup

	go func() {
		for i := 0; i < 200; i++ {
			ch <- i
		}
		close(ch)
	}()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go work(&wg, ch)
	}
	wg.Wait()
	fmt.Println("exit")
}

func work(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	for {
		time.Sleep(100 * time.Millisecond)
		select {
		case data, ok := <-ch:
			fmt.Println(data, ok)
		default:
			fmt.Println("not data")
		}

		_, ok := <-ch
		if !ok {
			break
		}
	}
}
