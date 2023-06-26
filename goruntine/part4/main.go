package main

import "fmt"

func worker(stop chan int) {

	for i := 0; i < 10; i++ {
		stop <- i
		//fmt.Println("干活....")
	}
	close(stop)
}

func main() {
	stop := make(chan int)
	go worker(stop)

	for i := range stop {
		fmt.Println(i)
	}

}
