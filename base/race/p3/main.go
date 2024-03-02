package main

func main() {
	watit := make(chan struct{})

	x := 0

	go func() {
		x++
		close(watit)
	}

	x++

	<-watit

	println(x)
}
