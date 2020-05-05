package main

import (
	"fmt"
	"time"
)

func main() {
	//slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	////s1 := slice[2:5]
	//s2 := slice[2:6:7]
	////
	////s2 = append(s2, 100)
	////s2 = append(s2, 200)
	////
	////s1[2] = 20
	//
	////fmt.Println(s1)
	//fmt.Println(s2)
	//fmt.Println(slice)
	//c()

	taskCh := make(chan int, 100)
	go worker(taskCh)

	// 塞任务
	for i := 0; i < 10; i++ {
		taskCh <- i
	}

	// 等待 1 小时
	select {
	case <-time.After(time.Second):
	}
}

func closeChannel()  {
	ch := make(chan int, 5)
	ch <- 18
	close(ch)
	x, ok := <-ch
	if ok {
		fmt.Println("received: ", x)
	}

	x, ok = <-ch
	if !ok {
		fmt.Println("channel closed, data invalid.")
	}
}



func worker(taskCh <-chan int) {
	const N = 5
	// 启动 5 个工作协程
	for i := 0; i < N; i++ {
		go func(j int) {
			for {
				task := <-taskCh
				fmt.Printf("finish task: %d by worker %d\n", task, j)
				time.Sleep(time.Second)
			}
		}(i)
	}
}

func c() {
	ticker := time.Tick(1 * time.Second)
	for {
		select {
		case <-ticker:
			// 执行定时任务
			fmt.Println("执行 1s 定时任务")

			return
		}

	}
}
