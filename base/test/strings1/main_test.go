package main

import (
	"strconv"
	"strings"
	"testing"
)

func BenchmarkString(b *testing.B) {
	var str string
	for i := 0; i < b.N; i++ {
		str += strconv.Itoa(i)
	}
}

func BenchmarkStringBuilder(b *testing.B) {
	var str strings.Builder
	for i := 0; i < b.N; i++ {
		str.WriteString(strconv.Itoa(i))
	}
}

// go test -bench=.
// 单个 go test -bench=BenchmarkString   或者   go test -bench=BenchmarkStringBuilder
