package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func requestHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 处理请求逻辑...
		// 这里只是简单地回复一个 "Hello, World!" 消息
		fmt.Fprint(w, "Hello, World!")
	}
}

func main() {
	// 创建一个 HTTP 服务器
	server := &http.Server{
		Addr:           ":8089",
		Handler:        requestHandler(),
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// 启动服务器
	go func() {
		if err := server.ListenAndServe(); err != nil {
			fmt.Printf("Server error: %v\n", err)
		}
	}()

	// 模拟发送请求，计算 QPS
	var wg sync.WaitGroup
	quit := make(chan struct{})
	qps := make(chan int, 1)
	requestCounter := 0 // 请求计数器
	startTime := time.Now()

	go func() {
		for {
			select {
			case <-quit:
				close(qps)
				return
			default:
				qps <- 0
				requestCounter++ // 递增请求计数器
				time.Sleep(time.Second / 500)
			}
		}
	}()

	go func() {
		for range qps {
			wg.Add(1)
			go func() {
				defer wg.Done()
				// 模拟请求处理
				// 处理逻辑...
				time.Sleep(10 * time.Millisecond)
			}()
		}
	}()

	// 等待一段时间
	time.Sleep(10 * time.Second)
	close(quit)

	wg.Wait()
	endTime := time.Now()
	elapsed := endTime.Sub(startTime).Seconds()

	// 计算实际 QPS
	actualQPS := float64(requestCounter) / elapsed
	fmt.Printf("Actual QPS: %.2f\n", actualQPS)
}
