package main

import (
	"fmt"
	"unsafe"
)

type F struct {
	Name string 
}

func NewF(name string) *F {
	return &F{Name: name}
}

func main() {
	a := Float64bits(11.1)
	fmt.Println(a)
}

func Float64bits(f float64) uint64 {
	return *(*uint64)(unsafe.Pointer(&f))
}

