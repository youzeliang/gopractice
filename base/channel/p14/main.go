package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func main() {
	urls := []string{
		"http://server1/resource",
		"http://server2/resource",
		"http://server3/resource",
		"http://server4/resource",
		"http://server5/resource",
	}

	results := make(chan string)
	var wg sync.WaitGroup

	// 启动协程
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			response, err := http.Get(url)
			if err == nil && response.StatusCode == http.StatusOK {
				body, _ := ioutil.ReadAll(response.Body)
				results <- string(body)
			}
		}(url)
	}

	// 等待获取结果
	go func() {
		wg.Wait()
		close(results)
	}()

	// 选择最先获取到的结果
	firstResult := <-results

	fmt.Println("First result:", firstResult)
}
