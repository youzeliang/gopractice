package main

import "fmt"

var done = make(chan bool)

var msg string

func aG() {
	msg = "hello"
	close(done)
	//done <- "fd"
}

func main() {

	go aG()

	a := <-done

	fmt.Println(a)

}
