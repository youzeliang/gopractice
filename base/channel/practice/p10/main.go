package main

import (
	"fmt"
	"time"
)

func main() {
	taskCh := make(chan int, 10)

	go worker(taskCh)

	for i := 0; i < 10; i++ {
		taskCh <- i
	}

	select {
	case <-time.After(time.Hour):
	}

}

func worker(taskCh chan int) {
	const N = 5
	// 启动 5 个工作协程
	for i := 0; i < N; i++ {
		go func(id int) {
			for {
				task := <-taskCh
				fmt.Printf("finish task: %d by worker %d\n", task, id)
				time.Sleep(time.Second)
			}
		}(i)
	}
}
