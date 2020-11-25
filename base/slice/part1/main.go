package main

import "fmt"

func main() {

	a := make([]int, 0, 10)
	a = append(a, 1, 2)
	a = append(a, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1)
	//FuncA(a)
	fmt.Println(a)
	//a[0] = 44
	//FuncA(a)
	//fmt.Println(a)
}

func FuncA(b []int) {
	b[0] = 19
}
