package main

import "fmt"

func worker(done chan bool) {
	fmt.Println("start working...")
	done <- true
	fmt.Println("end working...")
}

func main() {
	done := make(chan bool, 1)

	go worker(done)

	// 如果删了后、worker 函数可能还没有来得及打印 "end working..."。这就意味着主线程和 worker 函数将并行执行，互相不等待对方完成。
	<-done
}
