package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3, 4}
	newSlice := slice[0:3]
	fmt.Println("before modifying underlying array:")
	fmt.Println("slice: ", slice)
	fmt.Println("newSlice: ", newSlice)
	fmt.Println()

	newSlice[0] = 6
	newSlice = append(newSlice, 1, 2, 3, 4, 5, 6)
	fmt.Println("after modifying underlying array:")
	fmt.Println("slice: ", slice)
	fmt.Println("newSlice: ", newSlice)

}
