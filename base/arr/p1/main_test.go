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

// go test -bench . main_test.go

// for range 的确会稍劣于 for 循环，当然这其中包含了编译器级别优化的结果（通常是静态单赋值，或者 SSA 链接）。

// 关闭优化开关，再次运行压力测试。

// go test -c -gcflags '-N -l' . main_test.go

// 当没有编译器优化时，两种循环的性能都明显下降， for range 下降得更为明显，性能也更加比经典 for 循环差。
