package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//Ts()

	ffaa()
}

func ffaa() {


	var a = []int{6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6}

	sum := 0
	for i := 0; i < len(a); i++ {
		if a[i] == 6 {
			sum++
		}
	}

	fmt.Println(sum)
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

	var apr = struct {
		Name string
		Age  int
	}{
		Name: "fdsfds",
		Age:  10,
	}

	fmt.Println(apr)

}
