package main

import "time"

const (
	fmat = "2006-01-02 15:04:05"
)

// 当c被置为nil时,select会一直打印default
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
			if !ok {
				c = nil
			}
		}
	}
}
