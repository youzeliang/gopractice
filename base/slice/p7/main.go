package main

import "fmt"

func main() {
	foo := []int{0, 0, 0, 42, 100}
	bar := foo[1:4]
	bar = append(bar, 99)
	fmt.Println("foo:", foo) // foo: [0 0 0 42 99]
	fmt.Println("bar:", bar) // bar: [0 0 42 99]

	//bar 的 cap 容量会到原始切片的末尾，所以当前 bar 的 cap 长度为 4。

	//如果要解决这样的问题，其实可以在截取时指定容量：

}
