package main

import "fmt"

func main() {

	tes()
}

func tes() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(11111)
		}
	}()

	panic("fdfsd")
}
