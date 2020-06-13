package main

import (
	"fmt"
	"os"
	"time"
)

func foo() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	var bar = make(map[int]int)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		for {
			bar[1] = 1
		}
	}()
	for {
		bar[1] = 1
	}
}

func f1() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(11)
		}
	}()

	var bar = make(map[int]int)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(222)
			}
		}()

		for {
			bar[1] = 1
		}
	}()

	for {
		bar[1] = 1
	}
}

func main() {
	defer fmt.Println("defer main") // will this be called when panic?
	var user = os.Getenv("USER_")
	go func() {
		defer func() {
			fmt.Println("defer caller")
			if err := recover(); err != nil {
				fmt.Println("recover success.")
			}
		}()
		func() {
			defer func() {
				fmt.Println("defer here")
			}()

			if user == "" {
				panic("should set user env.")
			}
			fmt.Println("after panic")
		}()
	}()

	time.Sleep(1 * time.Second)
	fmt.Printf("get result")
}
