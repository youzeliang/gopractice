package main

import (
	"context"
	"net/http"
)

type key int

const (
	requestIDKey key = iota
)

func WithRequestId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// 从请求中提取请求ID和用户信息
		requestID := req.Header.Get("X-Request-ID")

		// 创建子 context，并添加一个请求 Id 的信息
		ctx := context.WithValue(req.Context(), requestIDKey, requestID)

		// 创建一个新的请求，设置新 ctx
		req = req.WithContext(ctx)

		// 将带有请求 ID 的上下文传递给下一个处理器
		next.ServeHTTP(rw, req)
	})
}
