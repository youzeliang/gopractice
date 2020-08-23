package main

import (
	"fmt"
)

func main() {

	ch := make(chan int)
	go rev(ch)
	ch <-10
}

func rev(c chan int)  {
	res := <- c
	fmt.Println(res)
}

