package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
)

func test3() {
	// 调用请求对象的SetQueryParams()，传入map[string]string，由resty来帮我们拼接

	//client.R().
	//	SetQueryParams(map[string]string{
	//		"name": "dj",
	//		"age": "18",
	//	}).
	//	Get(...)
}

func test4() {
	// resty还提供一种非常实用的设置路径参数接口

	//client.R().
	//	SetPathParams(map[string]string{
	//		"user": "dj",
	//	}).
	//	Get("/v1/users/{user}/details")

	// 注意，路径中的键需要用{}包起来。
}

func test1() {
	client := resty.New()

	resp, err := client.R().Get("https://baidu.com")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Response Info:")
	fmt.Println("Status Code:", resp.StatusCode())
	fmt.Println("Status:", resp.Status())
	fmt.Println("Proto:", resp.Proto())
	fmt.Println("Time:", resp.Time())
	fmt.Println("Received At:", resp.ReceivedAt())
	fmt.Println("Size:", resp.Size())
	fmt.Println("Headers:")
	for key, value := range resp.Header() {
		fmt.Println(key, "=", value)
	}
	fmt.Println("Cookies:")
	for i, cookie := range resp.Cookies() {
		fmt.Printf("cookie%d: name:%s value:%s\n", i, cookie.Name, cookie.Value)
	}
}

func main() {

	test()
}

type Library struct {
	Name   string
	Latest string
}

type Libraries struct {
	Results []*Library
}

func test() {
	client := resty.New()

	libraries := &Libraries{}
	client.R().SetResult(libraries).Get("https://api.cdnjs.com/libraries")
	fmt.Printf("%d libraries\n", len(libraries.Results))

	for _, lib := range libraries.Results {
		fmt.Println("first library:")
		fmt.Printf("name:%s latest:%s\n", lib.Name, lib.Latest)
		break
	}
}
