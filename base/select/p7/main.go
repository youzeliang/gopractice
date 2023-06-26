package main

// 如果一个 select 控制结构中包含一个 default 表达式，那么这个 select 并不会等待其它的 Channel 准备就绪，而是会非阻塞地读取或者写入数据：
func main() {
	ch := make(chan int)
	select {
	case i := <-ch:
		println(i)

	default:
		println("default")
	}
}
