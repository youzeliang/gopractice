package main

import (
	"context"
	"fmt"
	"time"
)
// http://litang.me/post/golang-server-design/
func longtimeCostFunc(ctx context.Context, c chan<- int) {
	for i := 0; i < 10; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("calculate interrupted")
			return
		case <-time.After(time.Second):
			fmt.Println("calculating...")
		}
	}
	c <- 1
	fmt.Println("calculate finished")
}

func main() {
	result := make(chan int, 1)
	ctx, cancel := context.WithCancel(context.Background())
	go longtimeCostFunc(ctx, result)

	select {
	case r := <-result:
		fmt.Println("longtimeCostFunc return:", r)
	case <-time.After(3 * time.Second):
		// notify worker goroutine to exit
		cancel()
		// handle timeout error
		fmt.Println("too long to wait for the result")
	}

	// blocking main goroutine exit
	time.Sleep(time.Minute)

	return
}
