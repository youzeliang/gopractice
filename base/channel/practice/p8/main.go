package main

// 控制并发数
var limit = make(chan int, 3)

func main() {
	// …………

	var work []func()
	for _, w := range work {
		go func() {
			limit <- 1
			w()
			<-limit
		}()
	}
	// …………
}
