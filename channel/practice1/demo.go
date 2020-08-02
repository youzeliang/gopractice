package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := make(chan string)

	// 创建producer()函数的并发goroutine
	go producer("ABCABC------AAA", channel)
	go producer("EEEEEDDDDDDD------MMMMMM", channel)
	// 数据消费函数
	customer(channel)

}

func producer(header string, channel chan<- string) {

	for {
		// 将随机数和字符串格式化为字符串发送给通道
		channel <- fmt.Sprintf("%s: %v", header, rand.Int31())
		time.Sleep(time.Second * 2)
	}
}

func customer(channel <-chan string) {

	for {
		message := <-channel
		fmt.Println(message)
	}
}
