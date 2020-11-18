package main

import "fmt"

func main() {

	fmt.Println(111)
}

func test1(base int) (func(int) int, func(int) int) {

	// 定义2个函数，并返回
	// 相加

	add := func(i int) int {
		base += i

		return base
	}

	del := func(i int) int {
		base -= i

		return base
	}

	return add, del

}
