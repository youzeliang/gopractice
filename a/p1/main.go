package main

type Animal interface {
	Speak() string
}
type Dog struct {
}

func (d Dog) Speak() string {
	return "Woof!"

}

func makeSomething() {

	var d = Dog{}
	d.Speak()
}

func main() {

	makeSomething()
}
