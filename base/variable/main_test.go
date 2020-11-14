package main

import (
	"fmt"
	"testing"
)

func BenchmarkTes(b *testing.B) {

	for i := 0; i < b.N; i++ {
		traversalString()
	}
}

// 遍历字符串
func traversalString() {
	s := "pprof.cn博客"
	//for i := 0; i < len(s); i++ { //byte
	//	fmt.Printf("%v(%c) ", s[i], s[i])
	//}
	fmt.Println()
	for _, r := range s { //rune

		_ = r
		//fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
}
