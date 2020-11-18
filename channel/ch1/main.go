package main

import "fmt"

func main() {

	///
	//ch := make(chan int, 1)
	//ch <- 10
	//fmt.Println("发送成功")

	tes1()

}

func tes1() {

	ch := make(chan int, 10)

	for i := 0; i < 10; i++ {
		ch <- i
	}

	fmt.Println(ch)

}
