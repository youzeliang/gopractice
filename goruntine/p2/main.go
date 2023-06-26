package main

import (
	"fmt"
	"runtime"
	"time"
)

const NUM = 10000

func cal() {
	for i := 0; i < NUM; i++ {
		runtime.Gosched()
	}
}

func main() {
	// 只设置一个 Processor
	runtime.GOMAXPROCS(1)
	start := time.Now().UnixNano()
	go cal()
	for i := 0; i < NUM; i++ {
		runtime.Gosched()
	}
	end := time.Now().UnixNano()
	fmt.Printf("total %vns per %vns", end-start, (end-start)/NUM)
}
