package main

import (
	"fmt"
	"sync"
)

type Object struct {
	value int
}

func NewObject() interface{} {
	return &Object{}
}

func main() {

	pool := sync.Pool{
		New: NewObject,
	}

	// 从池中获取对象
	obj := pool.Get().(*Object)

	// 设置对象的值
	obj.value = 1

	// 输出对象的值
	println(obj.value)

	// 放回池中
	pool.Put(obj)

	// 再次从池中获取对象
	obj = pool.Get().(*Object)
	fmt.Println(obj.value)
}
