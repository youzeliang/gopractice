package main

import (
	"fmt"
	"time"
)

func goroutineA(a <-chan int) {
	val := <-a
	fmt.Println("goroutine A received data: ", val)
	return
}

func goroutineB(b <-chan int) {
	val := <-b
	fmt.Println("goroutine B received data: ", val)
	return
}

func main() {
	//ch := make(chan int)
	//go goroutineA(ch)
	//go goroutineB(ch)
	//ch <- 3
	//time.Sleep(time.Second)
	//
	//ch1 := make(chan struct{})
	//fmt.Println(ch1)

	GoBirthTime := "2006-01-02 15:04:05"

	c := "2022-04-13 18:49:24"

	var data []string

	var en = "12"
	var b = "14"


	if en > b {

	}


	data = append(data, c)

	first, _ := time.ParseInLocation(GoBirthTime, data[0], time.Local)

	firstTime := first.Unix()

	for i := 0; i < len(data); i++ {
		newTime, _ := time.ParseInLocation(GoBirthTime, data[i], time.Local)
		if newTime.Unix() >= firstTime {
			firstTime = newTime.Unix()
		}
	}

	for i := 0; i < len(data); i++ {

		p, _ := time.ParseInLocation(GoBirthTime, data[i], time.Local)

		fmt.Println(firstTime, p.Unix())
	}

}

func sqrt(x int) int {

	var a int = x

	for a*a > x {
		a = (a + x/a) / 2
	}
	return a
}
