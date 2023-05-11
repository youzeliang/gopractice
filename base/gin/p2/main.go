package main

import (
	"fmt"
	"unsafe"
)

// 洋葱模型

type Context struct {
	handles []func(c *Context)
	index   int8 // 代表上面func的索引
}

func (this *Context) Use(f func(c *Context)) {
	this.handles = append(this.handles, f)
}

func (this *Context) Get(path string, f func(c *Context)) {
	this.handles = append(this.handles, f)
}

func (this *Context) Next() {
	this.index++
	this.handles[this.index](this)
}

func (this *Context) Run() {
	this.handles[0](this) // 执行第一个函数
}

func Middleware1() func(c *Context) {
	return func(c *Context) {
		fmt.Println("middleware start")
		c.Next()
		fmt.Println("middleware end")
	}
}

func Middleware2() func(c *Context) {
	return func(c *Context) {
		fmt.Println("middleware2 start")
		c.Next()
		fmt.Println("middleware2 end")
	}
}

func N(n int) []struct{} {
	return make([]struct{}, n)
}
func main() {

	c := &Context{}
	c.Use(Middleware1())
	c.Use(Middleware2())

	c.Get("/", func(c *Context) {
		fmt.Println("Get handle func ")
	},
	)

	c.Run()

	//var a []struct{}
	d := make([]struct{}, 24)
	cc := N(23)
	fmt.Println(unsafe.Sizeof(cc))

	fmt.Println(unsafe.Sizeof(d))
}
