package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}

func a(ch chan bool, i int) {

	fmt.Println("go func", i, "go count", runtime.NumGoroutine())
	<-ch

	wg.Done()
}

func main() {
	// 模拟用户需求go业务的数量

	tackCount := 100
	ch := make(chan bool, 4)

	for i := 0; i < tackCount; i++ {
		wg.Add(1)
		ch <- true
		go a(ch, i)

	}

	wg.Wait()
}
