package main

import "testing"

func RemoveByAppend(s []int, i int) []int {
	return append(s[:i], s[i+1:]...)
}

// disordered
func RemoveByReplace(s []int, i int) []int {
	l := len(s) - 1
	s[i] = s[l]
	return s[:l]
}

func BenchmarkRemoveAppend(b *testing.B) {
	b.ReportAllocs()

	index := 2
	for i := 0; i < b.N; i++ {
		s := []int{1, 2, 3, 4, 5, 6}
		s = RemoveByAppend(s, index)
	}
}
func BenchmarkRemoveByReplace(b *testing.B) {
	b.ReportAllocs()

	index := 2
	for i := 0; i < b.N; i++ {
		s := []int{1, 2, 3, 4, 5, 6}
		s = RemoveByReplace(s, index)
	}
}
