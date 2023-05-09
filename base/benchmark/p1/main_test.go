package main

import "testing"

func appendOne(num int) []int {
	var res []int
	for i := 0; i < num; i++ {
		res = append(res, i)
	}
	return res
}

func appendMany(num int) []int {
	res := make([]int, 0, num)
	for i := 0; i < num; i++ {
		res = append(res, i)
	}
	return res
}

func BenchmarkAppendOne(b *testing.B) {
	num := 10000
	for i := 0; i < b.N; i++ {
		_ = appendOne(num)
	}
}

func BenchmarkAppendMany(b *testing.B) {
	num := 10000
	for i := 0; i < b.N; i++ {
		_ = appendMany(num)
	}
}

//  go test -bench . -benchmem
// AppendMany每次进行操作进行1次的内存分配，每次内存分配，分配了81920 B, 每次操作耗时12241 ns，这些指标都好于AppendOne。
//一次性分配需要的内存大小，slice不需要在扩大底层数组时进行内存分配，旧的底层数据依然能够复用，这显然减少了GC的压力。
