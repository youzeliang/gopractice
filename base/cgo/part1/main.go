package main

import (
	"fmt"
)

func main() {
	//C.puts(C.CString("hello\n"))

	var ints []int = make([]int,0,5)
	ints = append(ints, 1)

	fmt.Println(ints)

	ca := a{}
	fmt.Println(ca)
}

type a struct {
	c int64

	b int
	cc int64
	aaa string
	d bool
}