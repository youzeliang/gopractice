package main

import (
	"fmt"
	"math/rand"
	"time"
)

//编写一个程序，使用扇出模式同时生成 100 个随机数。让每个 goroutine 生成一个随机数，并通过缓冲channel将该数字返回给主 goroutine。
//设置缓冲区channel的大小，以便永远不会发送阻塞。不要分配比您需要的更多的缓冲区。让主 goroutine 显示它收到的每个随机数，然后终止程序。

const (
	goroutines = 100
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	// Create the buffer channel with a buffer for
	// each goroutine to be created.
	values := make(chan int, goroutines)

	// Iterate and launch each goroutine.
	for gr := 0; gr < goroutines; gr++ {

		// Create an anonymous function for each goroutine that
		// generates a random number and sends it on the channel.
		go func() {
			values <- rand.Intn(1000)
		}()
	}

	// Create a variable to be used to track received messages.
	// Set the value to the number of goroutines created.
	wait := goroutines

	// Iterate receiving each value until they are all received.
	// Store them in a slice of ints.
	var nums []int
	for wait > 0 {
		nums = append(nums, <-values)
		wait--
	}

	// Print the values in our slice.
	fmt.Println(nums)
}
