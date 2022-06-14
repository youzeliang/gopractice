package main

import "fmt"

func main() {

	var a = []int{1, 2, 0, 3}

	if len(a) > 3 {
		for i := 1; i < len(a)-1; i++ {
			if a[i-1] != 0 && a[i+1] != 0 && a[i] == 0 {
				fmt.Println(a[i])
				fmt.Println(i)
			}
		}
	}
}
