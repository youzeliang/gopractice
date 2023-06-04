package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func inc() func() int {
	var a int = 1
	return func() int {
		a++
		return a
	}
}

func main() {
	//pos, neg := adder(), adder()
	//for i := 0; i < 10; i++ {
	//	fmt.Println(
	//		pos(i),
	//		neg(-2*i),
	//	)
	//}

	//m := inc()

	//fmt.Println(inc()())
	//fmt.Println(inc()())
	//fmt.Println(inc()())

	x := 1
	f := func() {
		fmt.Println(x)
	}

	x = 2
	x = 3
	f() // 3

	x = 4
}
