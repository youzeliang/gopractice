package main

import "fmt"

type student struct {
	name string
	age  int
}

func main() {
	m := make(map[string]*student)
	stus := []student{
		{name: "pprof.cn", age: 18},
		{name: "测试", age: 23},
		{name: "博客", age: 28},
	}

	//遍历整个slice数组，依次赋值给map
	for i := 0; i < len(stus); i++ {
		m[stus[i].name] = &stus[i]
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}

}
