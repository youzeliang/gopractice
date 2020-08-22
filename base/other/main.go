package main

import "fmt"

const MAX int = 3

func main() {



	/* nextNumber 为一个函数，函数 i 为 0 */
	//nextNumber := getSequence()
	//
	///* 调用 nextNumber 函数，i 变量自增 1 并返回 */
	//fmt.Println(nextNumber())
	//fmt.Println(nextNumber())
	//fmt.Println(nextNumber())
	//
	///* 创建新的函数 nextNumber1，并查看结果 */
	//nextNumber1 := getSequence()
	//fmt.Println(nextNumber1())
	//fmt.Println(nextNumber1())}

	//pos, neg := add(), add()
	//for i := 0; i < 10; i++ {
	//	fmt.Println(
	//		pos(i),
	//		neg(-2*i),
	//	)
	//}
}

func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func add() func(int) int {

	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

func A(i int) (func(int) int, func(int) int) {
	return func(m int) int {
			i += m
			return i
		}, func(n int) int {
			i -= n
			return i
		}
}

func n() {

	defer func(n string) {
		fmt.Println("Welcome", n)
	}("fdsfds")
}
