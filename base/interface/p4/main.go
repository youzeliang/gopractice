package main

import "io"

type myWriter struct {
}

func (m myWriter) Write(p []byte) (n int, err error) {
	panic("implement me")
}

//func (m myWriter) Write(p6 []byte) (n int, err error) {
//	panic("implement me")
//}

/*func (w myWriter) Write(p6 []byte) (n int, err error) {
	return
}*/

func main() {
	// 检查 *myWriter 类型是否实现了 io.Writer 接口
	var _ io.Writer = (*myWriter)(nil)

	// 检查 myWriter 类型是否实现了 io.Writer 接口
	var _ io.Writer = myWriter{}
}
