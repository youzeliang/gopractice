package main

import "time"

const (
	fmat = "2006-01-02 15:04:05"
)

// for 循环里面的被关闭通道会被重复读取,以下例子会无限打印
func main() {
	c := make(chan int)

	go func() {
		time.Sleep(time.Second * 1)
		c <- 1
		close(c)
	}()

	for {
		select {
		case x, ok := <-c:
			println("case1", time.Now().Format(fmat), x, ok)
			time.Sleep(time.Millisecond * 500)
		default:
			println("case2", time.Now().Format(fmat))
			time.Sleep(time.Millisecond * 500)

		}
	}
}
