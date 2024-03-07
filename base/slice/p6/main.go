package main

import (
	"fmt"
	"sync"
)

func main() {
	a := make([]int, 0)

	buffer := make(chan int, 3)

	go func() {
		for v := range buffer {
			a = append(a, v)
		}
	}()

	var wg sync.WaitGroup

	for i := 0; i < 10000; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			buffer <- i
		}(i)
	}

	wg.Wait()

	fmt.Println(len(a))

}

// 如果是用 append 的方式，那么会有数据竞争，因为append是不安全的

func main1() {
	a := make([]int, 0)
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(i int) {
			a = append(a, i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(len(a))
}
