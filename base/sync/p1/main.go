package main

import (
	"fmt"
	"sync"
)

var pool *sync.Pool

func init() {
	pool = &sync.Pool{
		New: func() interface{} {
			fmt.Println("create")
			return "hello"
		},
	}
}

func main() {
	// 从池中获取对象
	obj := pool.Get().(string)
	fmt.Println(obj)

	// 放回池中

	pool.Put(obj)

	// 再次从池中获取对象
	obj = pool.Get().(string)
	fmt.Println(obj)
}
