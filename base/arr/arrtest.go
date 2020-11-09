package main

import (
	"fmt"
	"math/rand"
)

func main() {
	Ts()
}

func Ts() {

	slice := []int{1, 2}
	//slice := array[0:2]
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	newSlice := append(slice, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13)
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	rand.Int()

	//m := make(map[string]string, 0)

}
