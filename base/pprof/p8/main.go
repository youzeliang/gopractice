package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"
)

func main() {
	go func() {
		ip := "0.0.0.0:6060"
		if err := http.ListenAndServe(ip, nil); err != nil {
			fmt.Printf("start pprof failed on %s\n", ip)
			os.Exit(1)
		}
	}()

	outCh := make(chan int)
	stime := time.Now()
	// 每1s个goroutine
	for {
		time.Sleep(1 * time.Second)
		go alloc(outCh)
		fmt.Printf("last: %dseconds\n", int(time.Now().Sub(stime).Seconds()))
	}
}

// alloc分配1M内存，goroutine会阻塞，不释放内存,导致泄漏
func alloc(outCh chan<- int) {
	buf := make([]byte, 1024*1024*1)
	_ = len(buf)
	fmt.Println("alloc make buffer done")

	outCh <- 0 // 37行
	fmt.Println("alloc finished")
}
