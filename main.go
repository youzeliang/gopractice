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
<<<<<<< HEAD
	fmt.Println(1111)
}
=======
	a := Float64bits(11.1)
	fmt.Println(a)
}

func Float64bits(f float64) uint64 {
	return *(*uint64)(unsafe.Pointer(&f))
}

>>>>>>> ae9288816f12c72a6ac09705300e487a2e3ef736
