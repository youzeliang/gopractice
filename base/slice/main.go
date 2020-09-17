package main

import (
	"fmt"
)

func main() {

	str := "我的"
	s := []byte(str) //中文字符需要用[]rune(str)

	fmt.Println(s)
}