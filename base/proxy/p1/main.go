package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", ProxyHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func ProxyHandler(w http.ResponseWriter, r *http.Request) {
	// 解析目标URL
	targetURL := r.URL.Scheme + "://" + r.URL.Host + r.URL.Path
	if r.URL.RawQuery != "" {
		targetURL += "?" + r.URL.RawQuery
	}

	// 创建新的请求对象
	targetReq, err := http.NewRequest(r.Method, targetURL, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 复制客户端请求的头部信息到新请求
	targetReq.Header = r.Header

	// 发送请求到目标服务器
	client := &http.Client{}
	targetResp, err := client.Do(targetReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer targetResp.Body.Close()

	// 复制目标服务器响应的头部信息到客户端响应
	for key, values := range targetResp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	// 将目标服务器响应的内容写入客户端连接
	w.WriteHeader(targetResp.StatusCode)
	io.Copy(w, targetResp.Body)
}
