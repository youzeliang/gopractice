package main

import (
	"fmt"
	"sort"
)



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

var ch chan int  = make(chan int)


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


	//n := 0
	//reply := &n
	//Multiply(10, 5, reply)
	//fmt.Println("Multiply:", &reply) // Multiply: 50

	testChan()
}

func testChan()  {
	var ch2 chan string
	ch2 = make(chan string)
	var ch3 chan string
	ch3 = make(chan string, 1)

	ch4 := make(chan float32)
	ch5 := make(chan float64, 2)
	fmt.Printf("无缓冲 局部变量 chan ch2 : %v\n", ch2)
	fmt.Printf("有缓冲 局部变量 chan ch3 : %v\n", ch3)
	fmt.Printf("无缓冲 局部变量 chan ch4 : %v\n", ch4)
	fmt.Printf("有缓冲 局部变量 chan ch5 : %v\n", ch5)
}

func TestSlice()  {

	m := map[string]string{"q": "q", "w": "w", "e": "e", "r": "r", "t": "t", "y": "y"}
	var slice []string
	for k, _ := range m {
		slice = append(slice, k)
	}
	fmt.Printf("clise string is : %v\n", slice)
	sort.Strings(slice)
	fmt.Printf("sorted slice string is : %v\n", slice)
	for _, v := range slice {
		fmt.Println(m[v])
	}



}

// this function changes reply:
func Multiply(a, b int, reply *int) {
	*reply = a * b
}