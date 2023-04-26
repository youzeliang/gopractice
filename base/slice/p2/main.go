package main

import "fmt"

func main() {
	foo := []int{0, 0, 0, 42, 100}
	bar := foo[1:4]
	fmt.Println(foo)
	bar[1] = 99
	fmt.Println("foo:", foo)
	fmt.Println("bar:", bar)
}
