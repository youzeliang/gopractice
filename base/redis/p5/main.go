package main

import "fmt"

func findNextElement(arr []int, a int) int {
	var c int
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == a {
			c = arr[i+1]
		}
	}

	fmt.Println(c)

	return c
}
func main() {

	var a = []int{1, 2, 3, 4, 5}
	fmt.Println(findNextElement(a, 6))
}
