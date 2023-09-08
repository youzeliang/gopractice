package main

import (
	"context"
	"fmt"
	"time"
)

// AsyncCall 使用通道来表示超时
func AsyncCall() {
	ctx := context.Background()
	done := make(chan struct{}, 1)

	go func(ctx context.Context) {
		// 发送HTTP请求
		done <- struct{}{}
	}(ctx)

	select {
	case <-done:
		fmt.Println("call successfully!!!")
		return
	case <-time.After(time.Duration(800 * time.Millisecond)):
		fmt.Println("timeout!!!")
		return
	}
}
