package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

// 队列生产与消费
func main() {
	var t1 = makeTask("adoJob", 1000)
	var t2 = makeTask("xs25Job", 500)
	//var tick = time.Tick(time.Millisecond * 1)
	for {
		select {
		case task := <-t1:
			log.Println(task)
		case task := <-t2:
			log.Println(task)
		//case <-tick:
		default:

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
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000))) //假设生产任务占用时间
			ch <- fmt.Sprintf("%s,生产数据 %d", queueName, i)
			i++
		}
	}()
	return ch
}
