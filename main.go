package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x + y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	//c := make(chan int)
	//quit := make(chan int)
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		fmt.Println(<-c)
	//	}
	//	quit <- 0
	//}()
	//fibonacci(c, quit)


	//fmt.Println("---------------------------")
	//
	//s := "good bye"
	//var p *string = &s
	//*p = "ciao"
	//fmt.Printf("Here is the pointer p: %p\n", p) // prints address
	//fmt.Printf("Here is the string *p: %s\n", *p) // prints string
	//fmt.Printf("Here is the string s: %s\n", s) // prints same string


	n := 0
	reply := &n
	Multiply(10, 5, reply)
	fmt.Println("Multiply:", &reply) // Multiply: 50
}

// this function changes reply:
func Multiply(a, b int, reply *int) {
	*reply = a * b
}