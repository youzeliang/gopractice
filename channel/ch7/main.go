package main

import (
	"fmt"
)

func main() {

	cacheCh:=make(chan int,5)
	cacheCh <- 2
	cacheCh <- 3
	fmt.Println("cacheCh容量为:",cap(cacheCh),",元素个数为：",len(cacheCh))
}
