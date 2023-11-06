package main

import (
	"fmt"
	"github.com/panjf2000/ants"
)

func main() {
	// Use the common pool
	for i := 0; i < 10; i++ {
		i := i
		err := ants.Submit(func() {
			fmt.Println(i)
		})
		if err != nil {
			return
		}
	}
	//time.Sleep(time.Second)
}
