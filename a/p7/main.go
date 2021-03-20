package main

import (
	"fmt"
)

func main() {
	chan1 := make(chan struct{}, 1)
	chan2 := make(chan struct{}, 1)
	chan_final := make(chan struct{}, 1)
	chan1 <- struct{}{}
	N := 100
	go func(N int) {
		for i := 1; i <= N; i += 2 {
			<-chan1
			fmt.Println(i)
			chan2 <- struct{}{}
		}
	}(N)
	go func(N int) {
		for i := 2; i <= N; i += 2 {
			<-chan2
			fmt.Println(i)
			if i == N {
				chan_final <- struct{}{}
			}
			chan1 <- struct{}{}

		}

	}(N)
	<-chan_final
}
