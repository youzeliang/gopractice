package main

import (
	"fmt"
	"net/http"
)

// Middleware 设计一个中间件，实现类似于gin框架可以建立多个中间件还可以一次性返回
type Middleware func(http.Handler) http.Handler

func MiddlewareChain(h http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		h = middleware(h)
	}
	return h
}

func Logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Logger Middleware")
		h.ServeHTTP(w, r)
	})
}

func Auth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Auth Middleware")
		h.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})

	// 使用 MiddlewareChain 函数组装多个中间件
	handler := MiddlewareChain(mux, Logger, Auth)

	http.ListenAndServe(":8080", handler)
}
