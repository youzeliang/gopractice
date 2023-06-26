package main

import "fmt"

type S struct{}

func main() {
	var x S
	y := &x
	c := *identity(y)
	fmt.Println(c)
}
func identity(z *S) *S {
	return z
}
