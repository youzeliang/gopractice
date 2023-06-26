package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, i := range s {
		sum += i
	}

	c <- sum

	b := <-c

	fmt.Println(b)

}

func t1() {

	// 这里我们定义了一个可以存储整数类型的带缓冲通道
	// 缓冲区大小为2
	ch := make(chan int, 2)

	// 因为 ch 是带缓冲的通道，我们可以同时发送两个数据
	// 而不用立刻需要去同步读取数据
	ch <- 1
	ch <- 2

	// 获取这两个数据
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func t2() {
	s := []int{1, 2, 3, 4, 5, 6}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c

	fmt.Println(x)
	fmt.Println(y)
}

func fibonacci(n int, c chan int) {

	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}

	close(c)
}

func main() {

	c := make(chan int, 11)
	go fibonacci(cap(c), c)

	// range 函数遍历每个从通道接收到的数据，因为 c 在发送完 10 个
	// 数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据
	// 之后就结束了。如果上面的 c 通道不关闭，那么 range 函数就不
	// 会结束，从而在接收第 11 个数据的时候就阻塞了。

	for i := range c {
		fmt.Println(i)
	}

}
