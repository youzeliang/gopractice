package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(1)

	ctx, stop := context.WithCancel(context.Background())
	go func() {
		defer wg.Done()
		watch(ctx, "watch")
	}()

	time.Sleep(5 * time.Second) //先让监控狗监控5秒
	stop()                      //发停止指令
	wg.Wait()
}

func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "copy that, stop")
			return
		default:
			fmt.Println(name, "watching……")
		}
		time.Sleep(1 * time.Second)
	}
}
