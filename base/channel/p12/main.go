package main

import (
	"fmt"
	"time"
)

func main() {
	leak2()
}

func leak2() {
	ch := make(chan int)

	// 消费者 g1
	go func() {
		for result := range ch {
			fmt.Printf("result: %d\n", result)
		}
	}()

	// 生产者 g2
	ch <- 1
	ch <- 2
	close(ch)
	time.Sleep(time.Second) // 模拟耗时
	fmt.Println("main goroutine g2 done...")
}
