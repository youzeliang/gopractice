package main

type Animal interface {
	Speak() string
}

type Dog struct {
}

func (d Dog) Speak() string {
	return "hello"
}

func makeSomething() {
	var d Animal
	d = Dog{}
	d.Speak()
}

func main() {

}
