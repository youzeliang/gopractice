package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

// 这里使用2个goroutine往n大小的通道中模拟任务生产。select中的case哪个可以读取则打印出数据，每隔5秒我们来看一下生产的消息还有多少没有被打印过。

func main() {
	var t1 = makeTask("adoJob", 1000)
	var t2 = makeTask("xs25Job", 500)
	var tick = time.Tick(time.Second * 5)
	for {
		select {
		case task := <-t1:
			log.Println(task)
		case task := <-t2:
			log.Println(task)
		case <-tick:
			log.Println(fmt.Sprintf("队列挤压数量t1:%v个，t2:%v个", len(t1), len(t2)))
		}
		time.Sleep(time.Second * 1)
	}
}

// 生产数据
func makeTask(queueName string, n int) chan string {
	ch := make(chan string, n)
	go func() {
		i := 1
		for {
			// 假设生产任务占用时间
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
			ch <- fmt.Sprintf("%s,生产数据 %d", queueName, i)
			i++
		}
	}()
	return ch
}
