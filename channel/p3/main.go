package main

import (
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan struct{}, 50)
	for i := 0; i < 10000; i++ {
		ch <- struct{}{}
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			log.Println(i)
			//time.Sleep(time.Second)
			<-ch
		}(i)
	}
	wg.Wait()
}

// 上例中创建了缓冲区大小为 3 的 channel，在没有被接收的情况下，至多发送 3 个消息则被阻塞。
// 开启协程前，调用ch <- struct{}{}，若缓存区满，则阻塞。协程任务结束，调用 <-ch 释放缓冲区。

func pr() {

	for i := 0; i < 10000; i++ {
		log.Println(i)
	}
}
