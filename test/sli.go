package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var z1 int = 321

		fmt.Println(unsafe.Sizeof(z1))


}
