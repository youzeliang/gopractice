package main

import (
	"fmt"
	"sync"
)

func main() {

	var mu sync.Mutex

	mu.Lock()

	go func() {
		fmt.Println(11)
	}()

	mu.Unlock()

}
