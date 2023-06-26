package main

func main() {
	for range N(4) {
	}
}

func N(n int) []struct{} {
	return make([]struct{}, n)
}
