package main

import "fmt"

func main() {

	// 将一个具体类型的值赋值给一个interface类型的变量的时候，就同时把类型和值都赋值给了interface里的两个指针。如果这个具体类型的值是nil的话，interface变量依然会存储对应的类型指针和值指针。
	var p *int = nil
	var i interface{} = p
	fmt.Printf("%v %+v is nil %v\n\n", i, p, i == nil)
}
