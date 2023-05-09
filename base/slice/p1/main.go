package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	shadow := a[1:3]
	shadow = append(shadow, 100)
	fmt.Println(shadow, a)
}
