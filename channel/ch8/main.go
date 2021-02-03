package main

import (
	"fmt"
	"sync"
	"time"
)

// 协程优雅的退出
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	stopCh := make(chan bool)

	go func() {
		defer wg.Done()
		watch(stopCh, "watch")
	}()

	time.Sleep(5 * time.Second)
	stopCh <- true
	wg.Wait()
}

func watch(stopCh chan bool, name string) {

	for {
		select {
		case <-stopCh:
			fmt.Println(name, "copy that, stop")
			return
		default:
			fmt.Println(name, "watching……")
		}
		time.Sleep(1 * time.Second)
	}
}
