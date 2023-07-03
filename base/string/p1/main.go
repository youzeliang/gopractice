package main

import (
	"fmt"
)

func main() {
	str := "Golang罗尔"
	fmt.Println(len(str))
	fmt.Println(len([]rune(str)))
}
