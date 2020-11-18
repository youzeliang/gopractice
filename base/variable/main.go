package main

import "fmt"

const (
	a = iota + 100
	b
	c
	d
	e
	f //3,4
)

func main() {

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println("-----------------")
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println("-----------------")
	fmt.Println(e)
	fmt.Println(f)
}
