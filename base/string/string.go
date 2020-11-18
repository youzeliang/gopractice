package main

import (
	"fmt"
	"unsafe"
)

func main() {
	changeString()
}

func changeString() {

	s1 := "hello"

	byte1 := []byte(s1)

	byte1[0] = 'e'

	fmt.Println(string(byte1))

	var a int32 = 43342

	println(unsafe.Sizeof(a))

}
