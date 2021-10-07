package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter = struct{
		sync.RWMutex
		m map[string]int
	}{m: make(map[string]int)}

	counter.RLock()
	n := counter.m["some_key"]
	counter.RUnlock()
	fmt.Println("some_key:", n)


	a :=[]string{"1"}
	res := make(map[string][]string)
	res[a[0]] = append(res[a[0]], a[1])


}
