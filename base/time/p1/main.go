package main

import (
	"fmt"
	"time"
)

//go:linkname nanotime1 runtime.nanotime1
func nanotime1() int64

func main() {
	defer func(begin int64) {
		cost := (nanotime1() - begin) / 1000 / 1000
		fmt.Printf("cost = %dms \n", cost)
	}(nanotime1())

	time.Sleep(time.Second)
}
