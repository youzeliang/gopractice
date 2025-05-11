package main

import (
	"fmt"
	"time"
)

func main() {
	leak1()
}

// 无缓冲通道。对于无缓冲通道，发送操作会阻塞，直到有接收方准备好接收数据。
func leak1() {
	ch := make(chan int)
	// g1
	go func() {
		time.Sleep(2 * time.Second) // 模拟 io 操作
		ch <- 100                   // 模拟返回结果
	}()

	// g2
	// 阻塞住，直到超时或返回
	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Println("timeout! exit...")
	case result := <-ch:
		fmt.Printf("result: %d\n", result)
	}
}
