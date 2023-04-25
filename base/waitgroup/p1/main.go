package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(10)

	var output []int
	for i := 0; i < 10; i++ {
		go func(i int) {
			defer wg.Done()
			output = append(output, i)
		}(i)
	}

	wg.Wait()
	for i := 0; i < len(output); i++ {
		fmt.Println(output[i])
	}
}
