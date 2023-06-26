package main

import (
	"sync"
	"testing"
)

func Test_setup(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup()
		})
	}
}

func BenchmarkName(b *testing.B) {
	ch := make(chan byte, 4096)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for range ch {
		}
	}()
	b.SetBytes(1)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch <- byte(i)
	}
	close(ch)
	wg.Wait()
}

func BenchmarkCopy(b *testing.B) {
	from := make([]byte, b.N)
	to := make([]byte, b.N)
	b.ReportAllocs()
	b.ResetTimer()
	b.SetBytes(1)
	copy(to, from)
}
