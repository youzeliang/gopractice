package main

import (
	"fmt"
	"sync"
)

func main() {

	a := make([]int, 0)

	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {

		wg.Add(1)

		go func(i int) {
			a = append(a, i)
			wg.Done()
		}(i)
	}

	wg.Wait()

	fmt.Println(len(a))

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
