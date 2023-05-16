package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}

func doBusiness(ch chan bool, i int) {

	fmt.Println("go func ", i, " goroutine count = ", runtime.NumGoroutine())

	<-ch

	wg.Done()
}

func main() {
	// 模拟用户需求go业务的数量
	taskCnt := math.MaxInt64

	ch := make(chan bool, 3)

	for i := 0; i < taskCnt; i++ {
		wg.Add(1)
		ch <- true
		go doBusiness(ch, i)
	}

	wg.Wait()
}
