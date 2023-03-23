package main

import (
	"fmt"
	"runtime"
)

func main() {
	num := runtime.NumCPU() //获取主机的逻辑CPU个数
	fmt.Println(num)
	runtime.GOMAXPROCS(num) //设置可同时执行的最大CPU数

}
