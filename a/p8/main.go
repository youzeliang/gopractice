package main

import (
	"fmt"
	//      "sync"
	"time"
)

var count = uint64(0)

//var l sync.Mutex

func add() {
	for {
		//              l.Lock()
		//              fmt.Println("Start ++")
		count++
		//              l.Unlock()
	}
}

func main() {
	go add()
	time.Sleep(2 * time.Second)
	fmt.Println("Count =", count)
}
