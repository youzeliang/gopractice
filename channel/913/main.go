package main

import "fmt"

func worker(ch chan struct{}) {
	<-ch
	fmt.Println("do something")
}

func main() {
	ch := make(chan struct{})
	go worker(ch)
	ch <- struct{}{}
	close(ch)
}

//有时候使用 channel 不需要发送任何的数据，只用来通知子协程（goroutine）执行任务，或只用来控制协程的并发。这种情况下，使用空结构体作为占位符就非常合适了。
