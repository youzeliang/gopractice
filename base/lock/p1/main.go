package main

import "fmt"

func main() {
	var temp = make(map[int]int, 5)
	temp[1] = 1
	temp[2] = 2
	temp[3] = 3
	temp[4] = 4

	for i := 0; i < 5; i++ {
		go func(map[int]int) {
			temp[4] = 10
		}(temp)
	}

	fmt.Println(temp)
}
