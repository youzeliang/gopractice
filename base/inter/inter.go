package main

import "fmt"

type Mather interface {
	Add(a, b int) int
	Sub(a, b int64) int64
}

type Adder struct {
	sss string
}

//go:noinline
func (adder Adder) Add(a, b int) int {
	return a + b
}

//go:noinline
func (adder Adder) Sub(a, b int64) int64 {
	return a - b
}

func main() {

	var j uint32
	var Eface interface{} // outsmart compiler (avoid static inference)

	i := uint64(42)
	Eface = i
	j = Eface.(uint32)
	fmt.Println(j)

}


var apr = struct {
	Name string
	Age  int
}{
	Name: "zhanglinpeng",
	Age: 13,
}


var a = struct {
	Name string
}{Name: "fds"}