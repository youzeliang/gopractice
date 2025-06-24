package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

// 全局缓存（故意不清理）
var globalCache = make(map[int][]byte)

// 模拟内存泄漏的函数
func causeMemoryLeak() {
	// 泄漏类型1: 向全局缓存不断添加数据且永不释放
	for i := 0; i < 100; i++ {
		// 分配1MB内存
		data := make([]byte, 1024*1024)
		globalCache[i] = data // 添加到全局map永不删除
	}

	// 泄漏类型2: 阻塞的goroutine (永不退出的goroutine)
	ch := make(chan int)
	go func() {
		// 从channel接收数据（但永远不会有数据发送）
		val := <-ch
		fmt.Println("永远不会执行:", val)
	}()

	// 泄漏类型3: 未重置的定时器
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for range ticker.C {
			// 持续执行任务...
		}
	}()

	fmt.Println("已创建内存泄漏场景")
}

func main() {
	fmt.Println("内存泄漏演示开始")

	go func() {
		http.ListenAndServe("localhost:6061", nil)
	}()

	causeMemoryLeak()

	// 保持程序运行以便观察内存增长
	select {}
}
