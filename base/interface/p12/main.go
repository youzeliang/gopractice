package main

import "io"

type myWriter struct {
	// ...
}

func (w *myWriter) Write(p []byte) (n int, err error) {
	// 实现具体的写入逻辑
	return len(p), nil
}

func main() {
	// 创建一个 myWriter 对象
	writer := &myWriter{}

	// 注册 myWriter 对象
	registerWriter(writer)
}

func registerWriter(w io.Writer) {
	// 在这里执行注册操作
}
