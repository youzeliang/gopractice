package main

import "testing"

func benchmarkFib(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Fib(n)
	}
}

func BenchmarkFib1(b *testing.B)  { benchmarkFib(b, 1) }
func BenchmarkFib2(b *testing.B)  { benchmarkFib(b, 2) }
func BenchmarkFib3(b *testing.B)  { benchmarkFib(b, 3) }
func BenchmarkFib10(b *testing.B) { benchmarkFib(b, 10) }
func BenchmarkFib20(b *testing.B) { benchmarkFib(b, 20) }
func BenchmarkFib40(b *testing.B) { benchmarkFib(b, 40) }

func TestFib(t *testing.T) {
	type args struct {
		n int
	}
	var tests = []struct {
		name string
		args args
		want int
	}{

		{name: "n=0", args: args{n: 0}, want: 0},
		{name: "n=1", args: args{n: 1}, want: 1},
		{name: "n=5", args: args{n: 5}, want: 5},
		{name: "n=10", args: args{n: 10}, want: 55},
		{name: "n=20", args: args{n: 20}, want: 6765},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fib(tt.args.n); got != tt.want {
				t.Errorf("Fib() = %v, want %v", got, tt.want)
			}
		})
	}
}
