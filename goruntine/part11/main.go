package main

import (
	"context"
	"fmt"
)

// Go语言带有内存自动回收的特性，因此内存一般不会泄漏。但是Goroutine确实存在泄漏的情况，同时泄漏的Goroutine引用的内存同样无法被回收。
func main() {
	ch := func() <-chan int {
		ch := make(chan int)
		go func() {
			for i := 0; ; i++ {
				ch <- i
			}
		}()

		return ch
	}()

	for v := range ch {
		fmt.Println(v)
		if v == 5 {
			break
		}
	}
}

//上面的程序中后台Goroutine向通道输入自然数序列，main()函数中输出序列。但是当break跳出for循环的时候，后台Goroutine就处于无法被回收的状态了。我们可以通过context包来避免这个问题：

func main1() {
	ctx, cannel := context.WithCancel(context.Background())
	ch := func(ctx context.Context) <-chan int {
		ch := make(chan int)
		go func() {
			for i := 0; ; i++ {
				select {
				case <-ctx.Done():
					return
				case ch <- i:

				}
			}
		}()

		return ch
	}(ctx)

	for v := range ch {
		fmt.Println(v)
		if v == 5 {
			cannel()
			break
		}
	}
}
