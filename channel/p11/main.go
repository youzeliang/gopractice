package main

import "fmt"

func main() {

	var a = []int{1, 1, 2, 3, 4, 5}

	cnts := map[int]int{}

	for i := 0; i < len(a); i++ {
		cnts[a[i]]++
	}

	fmt.Println(cnts)
}
