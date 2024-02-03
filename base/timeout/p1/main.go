package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

// 如果将 hardWork 内部panic,
func hardWork(job interface{}) error {
	time.Sleep(time.Minute)
	return nil
}

func requestWork(ctx context.Context, job interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	// 如果这里是没有缓冲的channel，那么会阻塞，因为没有goroutine来接收，最终打印出来的goroutine数量会是1001
	done := make(chan error, 1)
	go func() {
		done <- hardWork(job)
	}()
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

func main() {
	const total = 1000
	var wg sync.WaitGroup
	wg.Add(total)
	now := time.Now()
	for i := 0; i < total; i++ {
		go func() {
			defer wg.Done()
			requestWork(context.Background(), "any")
		}()
	}
	wg.Wait()
	fmt.Println("elapsed:", time.Since(now))
	time.Sleep(2 * time.Second)

	fmt.Println("number of goroutines:", runtime.NumGoroutine())

}
