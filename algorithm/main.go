package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now() // 记录开始时间

	time.Sleep(time.Second * 1)

	entTime := time.Now()
	elapsedTime := entTime.Sub(startTime).Seconds()

	fmt.Println(elapsedTime)
}
