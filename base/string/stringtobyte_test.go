package main

import (
	"strings"
	"testing"
)

var s = strings.Repeat("hello", 1024)

func testDefault() {
	a := []byte(s)
	_ = string(a)
}

func testUnsafe() {
	a := String2bytes(s)
	_ = Bytes2String(a)
}

func BenchmarkTestDefault(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testDefault()
	}
}

func BenchmarkTestUnsafe(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testUnsafe()
	}
}
