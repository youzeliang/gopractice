package main

func makeSomething() {
	a := make(chan *int, 1)
	b := 1
	a <- &b
}

func main() {

}
