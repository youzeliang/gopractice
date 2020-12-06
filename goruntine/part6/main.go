package main

import (
	"fmt"
	"sync"
)

var total struct {
	sync.Mutex
	value int
}

func main() {

	var wg sync.WaitGroup

	wg.Add(2)
	go worker(&wg)
	go worker(&wg)
	wg.Wait()
	fmt.Println(total.value)
}

func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		total.Lock()
		total.value += i
		total.Unlock()
	}
}
