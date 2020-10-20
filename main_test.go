package main

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestFloat64bits(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("监控退出，停止了...")
				return
			default:
				fmt.Println("goroutine监控中...")
				time.Sleep(2 * time.Second)
			}
		}
	}(ctx)

	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel()
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}

func shortest(str []string) string {

	if len(str) == 0 {
		return ""
	}

	start := str[0]

	for _, j := range str {
		if len(j) < len(start) {
			start = j
		}
	}

	return start
}
