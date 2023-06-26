package main

// 第一个goroutine是一个计数器，用于生成0、1、2、……形式的整数序列，然后通过channel将该整数序列发送给第二个goroutine；第二个goroutine是一个求平方的程序
// 对收到的每个整数求平方，然后将平方后的结果通过第二个channel发送给第三个goroutine；第三个goroutine是一个打印程序，打印收到的每个整数。
func counter(out chan<- int) {

	for i := 1; i < 10; i++ {
		out <- i
	}
	close(out)
}

func squarter(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		println(v)
	}
}

func main() {

	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)

	go squarter(squares, naturals)

	printer(squares)

}
