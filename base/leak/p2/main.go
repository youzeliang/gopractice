package main

import (
	"context"
)

// 这是有问题的，已下是
//func main() {
//	ch := func() <-chan int {
//		ch := make(chan int)
//		go func() {
//			for i := 0; ; i++ {
//				ch <- i
//			}
//		} ()
//		return ch
//	}()
//
//	for v := range ch {
//		fmt.Println(v)
//		if v == 5 {
//			break
//		}
//	}
//}

func main() {

	ctx, cancel := context.WithCancel(context.Background())

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

	for i := range ch {
		println(i)
		if i == 10 {
			cancel()
			break
		}
	}
}
