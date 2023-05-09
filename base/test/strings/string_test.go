package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSplit1(t *testing.T) { //
	got := Split("a:b:c", ":")         // 程序输出的结果
	want := []string{"a", "b", "c"}    // 期望的结果
	if !reflect.DeepEqual(want, got) { // 因为slice不能比较直接，借助反射包中的方法比较
		t.Errorf("excepted:%v, got:%v", want, got) // 测试失败输出错误提示
	}
}

//func TestMoreSplit(t *testing.T) {
//
//	got := Split("abcd", "bc")
//
//	want := []string{"a", "d"}
//	if !reflect.DeepEqual(want, got) {
//		t.Errorf("excepted:%v, got:%v", want, got)
//	}
//}

//func TestSplit(t *testing.T) {
//	type args struct {
//		s   string
//		sep string
//	}
//	tests := []struct {
//		name       string
//		args       args
//		wantResult []string
//	}{
//
//		{
//			name:       "6666666",
//			args:       args{"a:b:c", ":"},
//			wantResult: []string{"a", "b", "c"},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if gotResult := Split(tt.args.s, tt.args.sep); !reflect.DeepEqual(gotResult, tt.wantResult) {
//				t.Errorf("Split() = %v, want %v", gotResult, tt.wantResult)
//			}
//		})
//	}
//}

func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("枯藤老树昏鸦", "老")
	}
}

func BenchmarkSplitParallel(b *testing.B) {
	b.SetParallelism(1) // 设置使用的CPU数
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Split("枯藤老树昏鸦", "老")
		}
	})
}

func ExampleSplit() {

	fmt.Println(Split("a:b:c", ":"))
	fmt.Println(Split("aa1aa", "1"))
	// Output:
	// [a b c]
	// [aa aa]
}
