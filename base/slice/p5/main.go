package main

import "fmt"

func main() {
	// 创建一个切片
	slice1 := []int{1, 2, 3, 4, 5}

	// 创建第二个切片，它与第一个切片共享底层数组
	slice2 := slice1[1:4]

	// 对第一个切片进行切片扩容
	slice1 = append(slice1, 6)

	// 打印两个切片
	fmt.Println(slice1) // 输出 [1 2 3 4 5 6]
	fmt.Println(slice2) // 输出 [2 3 4]
}
