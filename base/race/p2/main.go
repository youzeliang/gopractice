package main

import "fmt"

func main() {
	getNumber()
}

func getNumber() int {

	var i = 0

	defer func() {
		fmt.Println(i)
	}()

	i = 2

	return i
}
