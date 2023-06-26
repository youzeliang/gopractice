package main

import "fmt"

type Person struct {
	age int
}

func (p Person) howOld() int {
	return p.age
}

func (p *Person) growUp() {
	p.age += 1
}

func main() {
	// rolle 是值类型
	rolle := Person{age: 18}

	// 值类型 调用接收者也是值类型的方法
	fmt.Println(rolle.howOld())

	// 值类型 调用接收者是指针类型的方法
	rolle.growUp()
	fmt.Println(rolle.howOld())

	// ----------------------

	// stefno 是指针类型
	stefno := &Person{age: 100}

	// 指针类型 调用接收者是值类型的方法
	fmt.Println(stefno.howOld())

	// 指针类型 调用接收者也是指针类型的方法
	stefno.growUp()
	fmt.Println(stefno.howOld())
}
