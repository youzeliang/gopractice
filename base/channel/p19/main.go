package main

import (
	"fmt"
	"math"
	"time"
)

func main() {

	maxCount := math.MaxInt64

	maxConcurrent := 10 // 最大并发数为10

	ch := make(chan struct{}, maxConcurrent) // 创建一个容量为10的channel

	for i := 0; i < maxCount; i++ {

		ch <- struct{}{} // 在channel中放入一个空结构体

		go func(i int) {

			// 做一些各种各样的业务逻辑处理

			fmt.Printf("go func num: %d\n", i)

			time.Sleep(time.Second)

			<-ch // 从channel中取出一个空结构体

		}(i)

	}

}
