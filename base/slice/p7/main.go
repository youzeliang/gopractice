package main

import "fmt"

func main() {
	foo := []int{0, 0, 0, 42, 100}
	bar := foo[1:4]
	bar = append(bar, 99)
	fmt.Println("foo:", foo) // foo: [0 0 0 42 99]
	fmt.Println("bar:", bar) // bar: [0 0 42 99]

}
