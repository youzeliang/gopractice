package main

func test() func() {
	a, b := 0, 1

	f := func() {
		a = 2
		b = 3
	}

	return f
}
func main() {

}
