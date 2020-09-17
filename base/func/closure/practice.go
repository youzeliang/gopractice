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

	for _, stu := range stus {
		m[stu.name] = &stu
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
}

func filter(s []student, f func(student) bool) []student {

	var r []student
	for _, c := range s {
		if f(c) == true {
			r = append(r, c)
		}
	}

	return r

}
