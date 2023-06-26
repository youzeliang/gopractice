package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)

func BenchmarkName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprint(rand.Int())
	}
}

func BenchmarkName1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strconv.Itoa(rand.Int())
	}
}
