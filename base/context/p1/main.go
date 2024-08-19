package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	before := time.Now()
	preCtx, _ := context.WithTimeout(ctx, 100*time.Millisecond)
	go func() {
		childCtx, _ := context.WithTimeout(preCtx, 300*time.Millisecond)
		select {
		case <-childCtx.Done():
			after := time.Now()
			fmt.Println("child during:", after.Sub(before).Milliseconds())
		}
	}()
	select {
	case <-preCtx.Done():
		after := time.Now()
		fmt.Println("pre during:", after.Sub(before).Milliseconds())
	}
}

// 举一个例子来说明一下 Context 中的级联退出。下面的代码中 childCtx 是 preCtx 的子 Context，其设置的超时时间为 300ms。
// 但是 preCtx 的超时时间为 100 ms，因此父 Context 退出后，子 Context 会立即退出，实际的等待时间只有 100ms。
