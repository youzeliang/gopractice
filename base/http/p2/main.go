package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	// 初始化反向代理服务
	proxy, err := NewProxy()
	if err != nil {
		panic(err)
	}
	// 所有请求都由ProxyRequestHandler函数进行处理
	http.HandleFunc("/", ProxyRequestHandler(proxy))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func NewProxy() (*httputil.ReverseProxy, error) {
	targetHost := "http://my-api-server.com"
	url, err := url.Parse(targetHost)
	if err != nil {
		return nil, err
	}

	proxy := httputil.NewSingleHostReverseProxy(url)
	return proxy, nil
}

// ProxyRequestHandler 使用代理处理HTTP请求
func ProxyRequestHandler(proxy *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	}
}

// 在 Go 语言中，实现反向代理非常简单，Go 语言标准库 httputil 中为我们提供了封装好的反向代理实现方式。下面是一个最简单的实现反向代理的例子。

// NewProxy() 借助 httputil.NewSingleHostReverseProxy 函数生成了一个反向代理服务器。NewSingleHostReverseProxy 函数的参数是实际的后端服务器地址。
// 如果后端有多个服务器，那么可以用一些策略来选择某一个合适的后端服务地址，从而实现负载均衡策略。可以看到，最核心的代码其实只有一行：
// proxy := httputil.NewSingleHostReverseProxy(url)
