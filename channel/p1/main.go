package main

import (
	"fmt"
	"time"
)

func main() {

	//worker()

	x := make(chan int, 1)
	go func() {
		x <- 1
	}()
	dd := <-x
	fmt.Println(dd)

}

func worker() {
	ticker := time.Tick(1 * time.Second)
	for {
		select {
		case <-ticker:
			// 执行定时任务
			fmt.Println("执行 1s 定时任务")
		}
	}
}
