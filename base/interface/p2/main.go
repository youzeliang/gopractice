package main

import "fmt"

func main() {
	golang := Go{}
	php := PHP{}

	sayHello(golang)
	sayHello(php)
}

type IGreeting interface {
	sayHello()
}

func sayHello(i IGreeting) {
	i.sayHello()
}

type Go struct{}

func (g Go) sayHello() {
	fmt.Println("Hi, I am GO!")
}

type PHP struct{}

func (p PHP) sayHello() {
	fmt.Println("Hi, I am PHP!")
}
