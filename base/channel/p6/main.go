package main

// 同步显示
var msg string
var done = make(chan bool)

func setup() {
	msg = "hello, world"
	done <- true
}
func main() {

	go setup()
	<-done
	println(msg)
}
