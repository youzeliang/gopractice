package main

import (
	"fmt"
	"time"
)

func main() {
	//var a interface{} = nil
	//var b *int = nil

	//isNil(a)
	//isNil(b)

	fmt.Println(3312)
	go func() {
		for i := 0; i < 10; i++ {
			a := (6 - i) / (6 - i)
			fmt.Println(a)
			fmt.Println(i)
			time.Sleep(time.Second)
		}
	}()

	time.Sleep(time.Second * 10)
	fmt.Println(8888)

	time.Sleep(10)
}

func isNil(x interface{}) {
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("non-empty interface")
}
