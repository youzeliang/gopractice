package main

import "fmt"

func main() {

	te()
}

func te() {

	var a = "aaaffd"

	func(n string) {
		fmt.Println(n)
	}(a)

	var c add = func(a int, b int) int {
		return a + b
	}

	fmt.Println(c(1, 1))

}

type add func(a int, b int) int

func callsim()  {

	s := sim(func(a, b int) {
		a  = 1
		b = 2
	})
	s(1,2)
}

func sim(func(a,b int)) func(c, d int) int {

	return func(a, b int) int {
		return a + b
	}
}
