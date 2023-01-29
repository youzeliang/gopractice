package main

import (
	"fmt"
	"math/rand"
)

func main() {
		s()
}

func s() {

	defer func() {
		if err := recover(); err != nil {
			go a()
		}
	}()
	a()

	select {

	}

}

func a() {
	ss := rand.Intn(1)

	fmt.Println(ss, "--------")
	if 2 == 2 {
		fmt.Println(331)
		panic("ffff")
	}

	fmt.Println("end")
}
