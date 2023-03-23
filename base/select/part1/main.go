package main

import (
	"fmt"
	"time"
)

func main() {
	result := make(chan string, 1)
	go func() {
		//模拟网络访问
		time.Sleep(1 * time.Second)
		result <- "服务端结果"
	}()
	select {
	case v := <-result:
		fmt.Println(v)
	case <-time.After(2 * time.Second):
		fmt.Println("网络访问超时了")
	}

	time.Sleep(time.Second * 5)
	fmt.Println(11)

}
