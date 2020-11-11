package main

import "fmt"

func main() {
	c := Cat{}
	fmt.Println(c.Say())

	d := Dog{}
	fmt.Println(d.Say())
}

type Cat struct {
}

type Dog struct {
}

func (c Cat) Say() string {
	return "666"
}

func (d Dog) Say() string {

	return "dog"
}
