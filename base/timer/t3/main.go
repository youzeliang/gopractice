package main

import "fmt"
import "time"

func main() {
	timeInfo()
}

func timeInfo() {

	ticker := time.NewTicker(time.Second) //每隔1分钟 执行一次程序
	var n = 0
	for t := range ticker.C {
		if n > 5 {
			ticker.Stop() //终止定时器继续执行
			break
		}
		fmt.Println(t)
		n++
	}
}
