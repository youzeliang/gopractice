package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := Float64bits(11.1)
	fmt.Println(a)
}

func Float64bits(f float64) uint64 {
	return *(*uint64)(unsafe.Pointer(&f))
}
