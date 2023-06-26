package main

import "time"

// select 在遇到两个 <-ch 同时响应时其实会随机选择一个 case 执行其中的表达式
func main() {
	ch := make(chan int)
	go func() {
		for range time.Tick(1 * time.Second) {
			ch <- 0
		}
	}()

	for {
		select {
		case <-ch:
			println("case1")

		case <-ch:
			println("case2")
		}
	}
}
