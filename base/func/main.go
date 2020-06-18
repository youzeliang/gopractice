package main

import "fmt"

func main() {

	// 匿名函数
	c := func(a, b int) int {
		fmt.Println(a, b)
		return a + b
	}(1, 2)

	fmt.Println(c)

}
