package main

import "fmt"

func test() func() {
	a, b := 0, 1

	f := func() {
		a = 2
		b = 3
	}
	fmt.Println(a, b)

	return f
}
func main() {

}
