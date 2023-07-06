package main

import (
	"sync"
	"testing"
)

var m = map[string]string{
	"key1": "value1",
	"key2": "value2",
	"key3": "value3",
	"key4": "value4",
	"key5": "value5",
}

func BenchmarkOriginal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var wg sync.WaitGroup
		wg.Add(len(m))
		for k, v := range m {
			k := k
			v := v
			go func() {
				_ = k
				_ = v
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkOptimized(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var wg sync.WaitGroup
		wg.Add(len(m))
		for k, v := range m {
			x := struct{ k, v string }{k, v}
			go func() {
				_ = x.k
				_ = x.v
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

// go test -bench=. benchmark_test.go
