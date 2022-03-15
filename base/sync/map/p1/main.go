package main

import (
	"fmt"
	"sync"
)

func main() {

	var a = []int{1, 2, 3}
	l := len(a)
	for m := 1; m < l; m++ {

		fmt.Println(a[m])
	}

	fmt.Println("-------------------------")

	list := map[string]interface{}{
		"name":          "田馥甄",
		"birthday":      "1983年3月30日",
		"age":           34,
		"hobby":         []string{"听音乐", "看电影", "电视", "xxx"},
		"constellation": "白羊座",
	}

	var m sync.Map
	for k, v := range list {
		m.Store(k, v)
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		m.Store("age", 22)
		m.LoadOrStore("tag", 8888)
		wg.Done()
	}()

	go func() {
		m.Delete("constellation")
		m.Store("age", 18)
		wg.Done()
	}()

	wg.Wait()

	m.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
}
