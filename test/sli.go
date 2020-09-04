package main

import "container/heap"

func smallestK(arr []int, k int) []int {
	intHeap := make(IntHeap, 0, len(arr))
	heap.Init(&intHeap)
	for i := 0; i < len(arr); i++ {
		heap.Push(&intHeap, arr[i])
		if len(intHeap) > k {
			heap.Pop(&intHeap)
		}
	}

	result := make([]int, 0, k)
	for i := 0; i < k; i++ {
		result = append(result, (heap.Pop(&intHeap)).(int))
	}

	return result
}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i int, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	length := len(*h)
	val := (*h)[length-1]
	*h = (*h)[:length-1]

	return val
}
