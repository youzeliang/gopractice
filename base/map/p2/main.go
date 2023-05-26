package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	mu   sync.Mutex
	data map[string]string
}

func (s *SafeMap) Get(string2 string) (string, bool) {

	s.mu.Lock()
	defer s.mu.Unlock()

	if v, ok := s.data[string2]; ok {
		return v, true
	}

	return "", false

}

func (s *SafeMap) Delete(string2 string) {

	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.data, string2)

}

func (s *SafeMap) Set(key string, value string) {

	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = value

}

func main() {

	sm := &SafeMap{data: make(map[string]string)}

	// 添加键值对
	sm.Set("key1", "value1")
	sm.Set("key2", "value2")

	// 从SafeMap中获取值
	value, ok := sm.Get("key1")
	if ok {
		fmt.Println(value)
	}

	// 删除键值对
	sm.Delete("key2")

	// 遍历SafeMap中的所有键值对
	for key, value := range sm.data {
		fmt.Println(key, value)
	}
}
