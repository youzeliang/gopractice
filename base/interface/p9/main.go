package main

import "fmt"

type IGreeting interface {
	sayHello()
}

func sayHello(i IGreeting) {
	i.sayHello()
}

type Go struct {
}

func (g Go) sayHello() {
	fmt.Println("Hi, i am Go")
}

type PHP struct {
}

func (p PHP) sayHello() {
	fmt.Println("Hi, i am php")
}

func main() {

	// 编译器在调用 sayHello() 函数时，会隐式地将 golang, php 对象转换成 IGreeting 类型
	golang := Go{}
	php := PHP{}

	sayHello(php)
	sayHello(golang)
}
