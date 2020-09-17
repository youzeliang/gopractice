package main

import "testing"

func BenchmarkName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice()
	}
}

func slice() []int {
	x := make([]int, 1024)
	for i := 0; i < len(x); i++ {
		x[i] = i
	}

	return x
}

func array() [1024]int {
	var x [1024]int
	for i := 0; i < len(x); i++ {
		x[i] = i
	}
	return x
}

func BenchmarkSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice()
	}
}

//  go test -bench . -benchmem -gcflags "-N -l"