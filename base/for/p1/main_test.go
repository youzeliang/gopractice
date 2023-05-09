package main

import "testing"

func BenchmarkClassicForLoopIntArray(b *testing.B) {
	b.ReportAllocs()
	var arr [100000]int
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(arr); j++ {
			arr[j] = j
		}
	}
}

func BenchmarkForRangeIntArray(b *testing.B) {
	b.ReportAllocs()
	var arr [100000]int
	for i := 0; i < b.N; i++ {
		for j, v := range arr {
			arr[j] = j
			_ = v
		}
	}
}
