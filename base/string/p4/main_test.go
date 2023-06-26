package main

import (
	"fmt"
	"testing"
)

// Good
func BenchmarkJoinStrWithOperator(b *testing.B) {
	s1, s2, s3 := "foo", "bar", "baz"
	for i := 0; i < b.N; i++ {
		_ = s1 + s2 + s3
	}
}

// Bad
func BenchmarkJoinStrWithSprintf(b *testing.B) {
	s1, s2, s3 := "foo", "bar", "baz"
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s%s%s", s1, s2, s3)
	}
}
