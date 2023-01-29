package main

import (
	"fmt"
	"time"
)

func test(n int, c chan int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		test(n,c)

		c <- n
	}()
	 test(n,c)

	panic(fmt.Sprintf("panic: %d", n))
}

func aa(){
	panic(312)
}

func main() {
	//c := make(chan int, 0)
	//
	//for i := 0; i < 10; i++ {
	//	go test(i, c)
	//}
	//
	//for i := 0; i < 10; i++ {
	//	<-c
	//}
	//
	//select {
	//
	//}
	go a()

	time.Sleep(10)
	select {

	}
}

func a()  {
	fmt.Println(1)
	go b()
}

func b()  {
	fmt.Println(2)
	go c()
}

func c()  {

	fmt.Println(3)
	go d()
	fmt.Println(66)
	time.Sleep(6)
	panic(331231)

}

func d()  {
	fmt.Println(4)
}
