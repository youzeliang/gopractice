package p2

import (
	"context"
	"fmt"
	"time"
)

func AsyncCall(callback func(bool)) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*800)
	defer cancel()

	done := make(chan bool)
	go func(ctx context.Context, done chan<- bool) {
		defer func() {
			// 在请求完成后关闭资源
			// 这里可以根据实际情况进行处理
			fmt.Println("HTTP请求完成，关闭资源")
			done <- true
		}()

		// 模拟发送HTTP请求
		time.Sleep(time.Millisecond * 600)

	}(ctx, done)

	select {
	case <-ctx.Done():
		fmt.Println("调用成功，处理完成")
		callback(true)
	case <-time.After(time.Millisecond * 400):
		fmt.Println("超时")
		callback(false)
	}

	<-done
}

func main() {
	AsyncCall(func(success bool) {
		if success {
			fmt.Println("回调函数：调用成功")
		} else {
			fmt.Println("回调函数：超时")
		}
	})
}
