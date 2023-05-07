package main

import (
	"fmt"
	"testing"
	"time"
)

func BenchmarkFoo(b *testing.B) {
	expensiveSetup()
	b.ResetTimer() // Reset the benchmark timer
	for i := 0; i < b.N; i++ {
		fmt.Println(123)
	}
}

func expensiveSetup() {
	time.Sleep(time.Second * 2)
}
