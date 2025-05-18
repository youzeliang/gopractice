package main

import (
	"fmt"
	"runtime"
)

func a(ch chan bool, i int) {
	fmt.Println("go func", i, "go routine", runtime.NumGoroutine())
	<-ch
}

func main() {
	tackCount := 100

	ch := make(chan bool, 4)

	for i := 0; i < tackCount; i++ {
		ch <- true
		a(ch, i)
	}
}
