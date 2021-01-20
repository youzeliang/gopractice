package main

import "fmt"

func main() {
	//假设值都为1，这里只赋值3个
	var arr = [102400]int{1, 1, 1}
	for i, n := range &arr {
		//just ignore i and n for simplify the example
		_ = i
		_ = n
	}

}

func prtMap(myMap map[int]*int) {
	for key, value := range myMap {
		fmt.Printf("map[%v]=%v\n", key, *value)
	}
}
