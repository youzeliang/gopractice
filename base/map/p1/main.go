package main

import (
	"fmt"
	"sync"
)

func main() {
	m := map[string]string{
		"1": "1",
	}

	var wg sync.WaitGroup
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			wg.Done()
			a := m["1"]
			fmt.Println(a)
		}()
	}

	wg.Wait()

}
