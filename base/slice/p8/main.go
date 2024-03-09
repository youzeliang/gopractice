package main

import (
	"fmt"
)

func main() {

	foo := []int{0, 0, 0, 42, 100}
	bar := foo[1:4]
	fmt.Println(bar)
	bar = append(bar, 99)
	//bar = append(bar, 99, 1, 2, 3, 4, 5)  // foo  [0 0 0 42 100] 因为已经跟原来的没有关系了
	fmt.Println("foo:", foo) // foo: [0 0 0 42 99]
	fmt.Println("bar:", bar) // bar: [0 0 42 99]

}
