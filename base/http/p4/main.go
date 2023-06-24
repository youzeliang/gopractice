package main

import (
	"context"
	"fmt"
	"time"
)

type Res struct {
	Name string `json:"name"`
	Nick string `json:"nick"`
	Call string `json:"call"`
}

func Total() {

	res := &Res{}

	ctx := context.Background()

	A(ctx, res)
	B(ctx, res)
	C(ctx, res)

	fmt.Println(res)
}

func main() {

	Total()
}

func A(ctx context.Context, res *Res) {
	// http.post()  模拟第三方接口调用
	time.Sleep(time.Second * 3)
	res.Name = "321"
}

func B(ctx context.Context, res *Res) {
	// http.post()  模拟第三方接口调用
	time.Sleep(time.Second * 4)
	res.Nick = "465"
}

func C(ctx context.Context, res *Res) {
	// http.post()  模拟第三方接口调用
	time.Sleep(time.Second * 6)
	res.Call = "789"
}
