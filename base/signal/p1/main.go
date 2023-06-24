package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {

	// 收到来自操作系统的 interrupt 信号时，它会打印所需的消息，然后退出。
	killSignal := make(chan os.Signal, 1)
	signal.Notify(killSignal, os.Interrupt)
	go func() {
		for {
			fmt.Println("Doing Work")
			time.Sleep(1 * time.Second)
		}
	}()
	<-killSignal
	fmt.Println("Thanks for using Golang!")
}
