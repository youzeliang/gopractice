package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	go Working(ctx)
	time.Sleep(2 * time.Second)
	cancelFunc()
	// 等待一段时间，以确保工作协程接收到取消信号并退出
	time.Sleep(1 * time.Second)
}

func Working(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("done...")
			return
		default:
			fmt.Println("ing...")
		}
	}
}
