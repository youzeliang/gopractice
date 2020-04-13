package main

import "fmt"

func main() {

	t1()
}

func t2() {
	var x []int
	c := 0
	for i := 0; i < 100000; i++ {
		x = append(x, i)
		if c != cap(x) {
			c = cap(x)
			fmt.Printf("index=%d cap=%d\n", i+1, c)
		}
	}
}

// 扩容量和类型有关
func t1() {
	e := []int32{1, 2, 3}
	fmt.Println("cap of e before:", cap(e))
	e = append(e, 4)
	fmt.Println("cap of e after:", cap(e))

	f := []int{1, 2, 3}
	fmt.Println("cap of f before:", cap(f))
	f = append(f, 4, 2, 4, 1, 1, 1, 1, 1, 1, 1, 4, 2, 4, 1, 1, 1, 1, 1, 1, 1, 4, 2, 4, 1, 1, 1, 1, 1, 1, 1, 4, 2, 4, 1, 1, 1, 1, 1, 1, 1, 4, 2, 4, 1, 1, 1, 1, 1, 1, 1, 4, 2, 4, 1, 1, 1, 1, 1, 1, 1, 4, 2, 4, 1, 1, 1, 1, 1, 1, 1, 4, 2, 4, 1, 1, 1, 1, 1, 1, 1, 4, 2, 4, 1, 1, 1, 1, 1, 1, 1, 4, 2, 4, 1, 1, 1, 1, 1, 1, 1, 4, 2, 4, 1, 1, 1, 1, 1, 1, 1)
	fmt.Println("cap of f after:", cap(f))

}
