package main

func makeSomething() {
	a := make([]int, 1000)

	for i := 0; i < len(a); i++ {
		a[i] = i
	}
}
func main() {

	makeSomething()
}
