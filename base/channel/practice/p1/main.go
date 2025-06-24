package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	share := make(chan int)

	var wg sync.WaitGroup

	go func() {
		goroutine("Bill", share)
		wg.Done()
	}()

	go func() {
		goroutine("Joan", share)
		wg.Done()
	}()

	share <- 1

	wg.Wait()

}

func goroutine(name string, share chan int) {
	for {
		value, ok := <-share
		if !ok {
			fmt.Println("Channel closed")
			return
		}

		fmt.Printf("Goroutine %s Inc %d\n", name, value)
		if value == 10 {
			close(share)
			fmt.Printf("Goroutine %s Down\n", name)
			os.Exit(0) // Exit the program immediately
		}
		share <- value + 1

	}

}

//
//import (
//	"fmt"
//	"sync"
//)
//
//// 编写一个程序，其中两个 goroutine 来回传递一个整数十次。当每个 goroutine 接收到整数时打印。每次通过整数都增加。一旦整数等于 10，立刻终止程序。
//
//func main() {
//
//	share := make(chan int)
//
//	// Create the WaitGroup and add a count
//	// of two, one for each goroutine.
//	var wg sync.WaitGroup
//	wg.Add(2)
//
//	// Launch two goroutines.
//	go func() {
//		goroutine("Bill", share)
//		wg.Done()
//	}()
//
//	go func() {
//		goroutine("Joan", share)
//		wg.Done()
//	}()
//
//	share <- 1
//
//	// Wait for the program to finish.
//	wg.Wait()
//}
//
//// goroutine simulates sharing a value.
//func goroutine(name string, share chan int) {
//	for {
//
//		// Wait to receive a value.
//		value, ok := <-share
//		if !ok {
//
//			// If the channel was closed, return.
//			fmt.Printf("Goroutine %s Down\n", name)
//			return
//		}
//
//		// Display the value.
//		fmt.Printf("Goroutine %s Inc %d\n", name, value)
//
//		// Terminate when the value is 10.
//		if value == 10 {
//			close(share)
//			fmt.Printf("Goroutine %s Down\n", name)
//			return
//		}
//
//		// Increment the value and send it
//		// over the channel.
//		share <- value + 1
//	}
//}
