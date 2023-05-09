package main

import (
<<<<<<< HEAD
	"reflect"
	"strings"
	"testing"
)

func TestSplit(t *testing.T) { // 测试函数名必须以Test开头，必须接收一个*testing.T类型参数
	got := strings.Split("a:b:c", ":") // 程序输出的结果
	want := []string{"a", "b", "c"}    // 期望的结果
	if !reflect.DeepEqual(want, got) { // 因为slice不能比较直接，借助反射包中的方法比较
		t.Errorf("excepted:%v, got:%v", want, got) // 测试失败输出错误提示
	}

}

func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Split("枯藤老树昏鸦", "老")
	}
=======
	"fmt"
	"testing"
	"time"
)

func BenchmarkFoo(b *testing.B) {
	expensiveSetup()
	b.ResetTimer() // Reset the benchmark timer
	for i := 0; i < b.N; i++ {
		fmt.Println(123)
	}
}

func expensiveSetup() {
	time.Sleep(time.Second * 2)
>>>>>>> 9dad412f07f955fe7cf7c2792221d5e58f71b91e
}
