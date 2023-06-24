package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Res struct {
	Name string `json:"name"`
	Nick string `json:"nick"`
	Call string `json:"call"`
}

func main() {
	Total()
}

func Total() {
	res := &Res{}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		A(ctx, res)
	}()

	go func() {
		defer wg.Done()
		B(ctx, res)
	}()

	go func() {
		defer wg.Done()
		C(ctx, res)
	}()

	wg.Wait()
	fmt.Println(res)
}

func A(ctx context.Context, res *Res) {
	apiCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// 调用第三方接口
	result, err := CallAPI(apiCtx)
	if err != nil {
		fmt.Println("A: API call error:", err)
		return
	}

	fmt.Println(res)

	select {
	case <-ctx.Done():
		return
	case <-time.After(3 * time.Second):
		res.Name = result
	}
}

func B(ctx context.Context, res *Res) {
	apiCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// 调用第三方接口
	result, err := CallAPI(apiCtx)
	if err != nil {
		fmt.Println("B: API call error:", err)
		return
	}

	select {
	case <-ctx.Done():
		return
	case <-time.After(6 * time.Second):
		res.Nick = result
	}
}

func C(ctx context.Context, res *Res) {
	apiCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// 调用第三方接口
	result, err := CallAPI(apiCtx)
	if err != nil {
		fmt.Println("C: API call error:", err)
		return
	}

	select {
	case <-ctx.Done():
		return
	case <-time.After(6 * time.Second):
		res.Call = result
	}

	fmt.Println(res)
}

func CallAPI(ctx context.Context) (string, error) {
	// 模拟第三方接口调用
	time.Sleep(1 * time.Second)
	return "API Result", nil
}
