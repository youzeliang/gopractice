package main

import "fmt"

func main() {
	var m1 map[string]int
	var m2 = make(map[string]int)
	if m1 != nil {
		println("m1 not nil")
	} else {
		println("m1 nil")
	}
	if m2 != nil {
		println("m2 not nil")
	} else {
		println("m2 nil")
	}

	println(len(m1), len(m2))

	contentList := make([]int, 4)

	fmt.Println(contentList)

}
