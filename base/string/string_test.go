package main

import (
	"bytes"
	"testing"
)

func passedByValue(foo string) {
	foo.C = "1"
}

func passedByPointer(bar *string) {
	bar.C = "1"
}

// 值传递
func Benchmark_PassedByValue(b *testing.B) {
	var val string
	str := bytes.Buffer{}
	// 这里为了构建一个大值进行传递，小值因为copy代价太小性能差距不明显。
	for i := 0; i < 10000000; i++ {
		str.Write([]byte("====="))
	}
	val.C = str.String()

	for i := 0; i < b.N; i++ {
		passedByValue(val)
	}
}

// 指针传递
func Benchmark_PassedByPointer(b *testing.B) {
	var val = new(Value)
	str := bytes.Buffer{}
	for i := 0; i < 10000000; i++ {
		str.Write([]byte("====="))
	}
	val.C = str.String()
	for i := 0; i < b.N; i++ {
		passedByPointer(val)
	}
}
