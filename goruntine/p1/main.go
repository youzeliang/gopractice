package main

import (
	"fmt"
	"sync"
)

func main() {
	g()
}

func g() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int, wg sync.WaitGroup) {
			fmt.Println(i)
			defer wg.Done()
		}(i, wg)
	}
	wg.Wait()
	fmt.Println("finished")

	//time.Sleep(time.Second*10)
}

// waitGroup 的add 要在goroutine之前执行
