package main

import (
	"context"
	"fmt"
	"sync"
)

var (
	reqChannels     []chan RequestInfo
	CtxParent       context.Context
	ParentCancel    context.CancelFunc
	WaitParentGroup *sync.WaitGroup
)

//请求消息
type RequestInfo struct {
	Context    context.Context
	ReqData    interface{}
}

func init() {
	channelNum := 128
	channelSize := 1024
	reqChannels = make([]chan RequestInfo, channelNum)
	for i := 0; i < channelNum; i++ {
		reqChannels[i] = make(chan RequestInfo, channelSize)
		WaitParentGroup.Add(1)
		go channelHandler(CtxParent, i)
	}
}

func channelHandler(ctx context.Context, chanIdx int) int{

	defer WaitParentGroup.Done()

	for {
		select {
		case request := <-reqChannels[chanIdx]:

		case <-ctx.Done():
			return
		}
	}
}


func main() {

	fmt.Println(int64(-1))
}

//func main() {
//	//http://127.0.0.1:8000/go
//	// 单独写回调函数
//	http.HandleFunc("/go", myHandler)
//	//http.HandleFunc("/ungo",myHandler2 )
//	// addr：监听的地址
//	// handler：回调函数
//	http.ListenAndServe("127.0.0.1:8000", nil)
//}
//
//// handler函数
//func myHandler(w http.ResponseWriter, r *http.Request) {
//	fmt.Println(r.RemoteAddr, "连接成功")
//	// 请求方式：GET POST DELETE PUT UPDATE
//	fmt.Println("method:", r.Method)
//	// /go
//	fmt.Println("url:", r.URL.Path)
//	fmt.Println("header:", r.Header)
//	fmt.Println("body:", r.Body)
//	fmt.Println("body:", r.TLS)
//	// 回复
//	w.Write([]byte("www.5lmh.com"))
//}
