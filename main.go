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
	fmt.Println(1111)
}

func Float64bits(f float64) uint64 {
	return *(*uint64)(unsafe.Pointer(&f))
}
