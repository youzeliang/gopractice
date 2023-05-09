package main

import "testing"

type Sumifier interface{ Add(a, b int32) int32 }

type Sumer struct{ id int32 }

func (math Sumer) Add(a, b int32) int32 { return a + b }

type SumerPointer struct{ id int32 }

func (math *SumerPointer) Add(a, b int32) int32 { return a + b }

func BenchmarkDirect(b *testing.B) {
	adder := Sumer{id: 6754}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		adder.Add(10, 12)
	}
}

func BenchmarkInterface(b *testing.B) {
	adder := Sumer{id: 6754}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Sumifier(adder).Add(10, 12)
	}
}

func BenchmarkInterfacePointer(b *testing.B) {
	adder := &SumerPointer{id: 6754}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Sumifier(adder).Add(10, 12)
	}
}

// 用一个简单的 Benchmark 测试来说明一下。在下面这个例子中，BenchmarkDirect 测试直接调用调用的开销。BenchmarkInterface 测试进行接口调用的开销，
// 但其函数接收者是一个非指针。BenchmarkInterfacePointer 也是测试接口调用的开销，但其函数接收者是一个指针。
// 在 Benchmark 测试中，静止编译器的优化和内联汇编，避免这两种因素对耗时产生的影响。测试结果如下。可以看到直接函数调用的速度最快，为 xxx ns/op， 方法接收者为指针的接口调用和函数调用的速度类似，为 xxx ns/op， 方法接收者为非指针的接口调用却慢了数倍，为 xxxx ns/op。
// 方法接收者为非指针的接口调用速度之所以很慢是受到了内存拷贝的影响。由于接口中存储了数据的指针，而函数调用的是非指针，因此数据会从对堆内存拷贝到栈内存，让调用速度变慢。
