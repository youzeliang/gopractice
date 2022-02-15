package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	waitTime()
}

func waitTime() {
	var w = sync.WaitGroup{}
	var ch = make(chan bool)
	w.Add(2)
	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("我2秒")
		w.Done()
	}()
	go func() {
		time.Sleep(time.Second * 6)
		fmt.Println("我6秒")
		w.Done()
	}()
	go func() {
		w.Wait()
		ch <- false
	}()

	select {
	case <-time.After(time.Second * 5):
		fmt.Println("我超时了")
	case <-ch:
		fmt.Println("我结束了")
	}
}
