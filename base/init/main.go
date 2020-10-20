package main

import "fmt"

func main() {
	//fmt.Println("start")

	//a := 4
	//a <<= 1

	var a = []int{1, 2, 3, 4, 5}

	res := 0
	for i := range a {

		res ^= a[i]

		fmt.Println(res)
	}

	fmt.Println(res)


	fmt.Println(2^0)

}
