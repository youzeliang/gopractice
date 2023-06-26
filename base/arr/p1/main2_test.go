package main

import "testing"

type U5 struct {
	a, b, c, d, e int
}
type U4 struct {
	a, b, c, d int
}
type U3 struct {
	b, c, d int
}
type U2 struct {
	c, d int
}
type U1 struct {
	d int
}

func BenchmarkClassicForLoopLargeStructArrayU5(b *testing.B) {
	b.ReportAllocs()
	var arr [100000]U5
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(arr)-1; j++ {
			arr[j].d = j
		}
	}
}
func BenchmarkClassicForLoopLargeStructArrayU4(b *testing.B) {
	b.ReportAllocs()
	var arr [100000]U4
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(arr)-1; j++ {
			arr[j].d = j
		}
	}
}
func BenchmarkClassicForLoopLargeStructArrayU3(b *testing.B) {
	b.ReportAllocs()
	var arr [100000]U3
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(arr)-1; j++ {
			arr[j].d = j
		}
	}
}
func BenchmarkClassicForLoopLargeStructArrayU2(b *testing.B) {
	b.ReportAllocs()
	var arr [100000]U2
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(arr)-1; j++ {
			arr[j].d = j
		}
	}
}

func BenchmarkClassicForLoopLargeStructArrayU1(b *testing.B) {
	b.ReportAllocs()
	var arr [100000]U1
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(arr)-1; j++ {
			arr[j].d = j
		}
	}
}

func BenchmarkForRangeLargeStructArrayU5(b *testing.B) {
	b.ReportAllocs()
	var arr [100000]U5
	for i := 0; i < b.N; i++ {
		for j, v := range arr {
			arr[j].d = j
			_ = v
		}
	}
}
func BenchmarkForRangeLargeStructArrayU4(b *testing.B) {
	b.ReportAllocs()
	var arr [100000]U4
	for i := 0; i < b.N; i++ {
		for j, v := range arr {
			arr[j].d = j
			_ = v
		}
	}
}

func BenchmarkForRangeLargeStructArrayU3(b *testing.B) {
	b.ReportAllocs()
	var arr [100000]U3
	for i := 0; i < b.N; i++ {
		for j, v := range arr {
			arr[j].d = j
			_ = v
		}
	}
}
func BenchmarkForRangeLargeStructArrayU2(b *testing.B) {
	b.ReportAllocs()
	var arr [100000]U2
	for i := 0; i < b.N; i++ {
		for j, v := range arr {
			arr[j].d = j
			_ = v
		}
	}
}
func BenchmarkForRangeLargeStructArrayU1(b *testing.B) {
	b.ReportAllocs()
	var arr [100000]U1
	for i := 0; i < b.N; i++ {
		for j, v := range arr {
			arr[j].d = j
			_ = v
		}
	}
}

// 在这个例子中，定义了 5 种类型的结构体：U1~U5，它们的区别在于包含的 int 类型字段的数量。

// go test -bench . main2_test.go

// 现象 不管是什么类型的结构体元素数组，经典的 for 循环遍历的性能比较一致，但是 for range 的遍历性能会随着结构字段数量的增加而降低。

// 带着疑惑，发现了一个与这个问题相关的 issue：cmd/compile: optimize large structs：https://github.com/golang/go/issues/24416。这个 issue 大致是说：如果一个结构体类型有超过一定数量的字段（或一些其他条件），就会将该类型视为 unSSAable。如果 SSA 不可行，那么就无法通过 SSA 优化，这也是造成上述基准测试结果的重要原因。

// 结论 对于遍历大数组而言， for 循环能比 for range 循环更高效与稳定，这一点在数组元素为结构体类型更加明显。
//
//另外，由于在 Go 中切片的底层都是通过数组来存储数据，尽管有 for range 的副本复制问题，但是切片副本指向的底层数组与原切片是一致的。这意味着，当我们将数组通过切片代替后，不管是通过 for range 或者 for 循环均能得到一致的稳定的遍历性能。

// 参考 https://mp.weixin.qq.com/s/dHjGn3gwxnYqlrsUkgc0dg
