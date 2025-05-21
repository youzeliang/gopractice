package main

import (
	"container/heap"
	"fmt"
	"math/rand"
	"time"
)

// 定义最小堆类型（基于整型切片）
type MinHeap []int

// 实现 heap.Interface 接口的方法
func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] } // 最小堆的比较逻辑
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 查找前100大值的函数
func findTop100(data []int) []int {
	h := &MinHeap{}
	heap.Init(h)

	// 填充前100个元素
	for _, num := range data[:100] {
		heap.Push(h, num)
	}

	// 遍历剩余元素
	for _, num := range data[100:] {
		if num > (*h)[0] { // 当前元素比堆顶大
			heap.Pop(h)       // 移除堆顶最小值
			heap.Push(h, num) // 插入新值
		}
	}

	// 提取堆中元素并排序（降序）
	result := make([]int, h.Len())
	for i := len(result) - 1; i >= 0; i-- {
		result[i] = heap.Pop(h).(int)
	}
	return result
}

func main() {
	// 生成测试数据（1亿个随机整数）
	rand.Seed(time.Now().UnixNano())
	data := make([]int, 1e8) // 1亿数据量（根据内存调整测试规模）
	for i := range data {
		data[i] = rand.Intn(1e9) // 生成0~1e9之间的随机数
	}

	// 查找前100大值
	start := time.Now()
	top100 := findTop100(data)
	elapsed := time.Since(start)

	fmt.Printf("Top 100 values: %v\n", top100)
	fmt.Printf("Execution time: %s\n", elapsed)
}
