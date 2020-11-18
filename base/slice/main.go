package main

import (
	"fmt"
)

func main() {

	//sl()
	panic(321321)

}

func sss() {
	str := "我的"
	s := []byte(str) //中文字符需要用[]rune(str)

	fmt.Println(s)
}

func sl() {

	var a = [3]int{1, 2, 3}

	for i := 0; i < len(a); i++ {
		fmt.Println(i)
	}

}
