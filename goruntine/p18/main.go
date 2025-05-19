package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}

func a(ch chan int) {

	for t := range ch {
		fmt.Println("go task = ", t, ", goroutine count = ", runtime.NumGoroutine())
		wg.Done()
	}
}

func sendTask(task int, ch chan int) {
	wg.Add(1)
	ch <- task
}

func main() {
	// 无buffer channel
	ch := make(chan int)

	goCnt := 3

	for i := 0; i < goCnt; i++ {
		// 启动go
		go a(ch)
	}

	tackCnt := math.MaxInt64
	for i := 0; i < tackCnt; i++ {
		sendTask(i, ch)
	}
}
