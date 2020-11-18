package main

import "testing"

func DoDefer(key, value string) {
	defer func(key, value string) {
		_ = key + value
	}(key, value)
}

func DoNotDefer(key, value string) {
	_ = key + value
}

func BenchmarkDoDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DoDefer("煎鱼", "https://github.com/EDDYCJY/blog")
	}
}

func BenchmarkDoNotDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DoNotDefer("煎鱼", "https://github.com/EDDYCJY/blog")
	}
}
