package main

import (
	"log"
	"runtime"
	"time"
)

func test() {
	//slice 会动态扩容，用slice来做堆内存申请
	container := make([]int, 8)

	log.Println(" ===> loop begin.")
	for i := 0; i < 32*1000*1000; i++ {
		container = append(container, i)
	}
	log.Println(" ===> loop end.")
}

func main() {
	log.Println("Start.")

	test()

	log.Println("force gc.")
	runtime.GC() //强制调用gc回收

	log.Println("Done.")

	time.Sleep(3600 * time.Second) //睡眠，保持程序不退出
}

func find(nums []int) int {

	n := len(nums)
	if n == 0 {
		return 0
	}

	// i1 是 nums 中第 1 大的数的索引号
	// i2 是 nums 中第 2 大的数的索引号
	i1, i2 := 0, 1
	if nums[i1] < nums[i2] {
		i1, i2 = i2, i1
	}

	for i := 2; i < n; i++ {
		if nums[i1] < nums[i] {
			i2, i1 = i1, i
		} else if nums[i2] < nums[i] {
			i2 = i
		}
	}
	return 0
}

func container(nums []int, k int) bool {
	if k < 0 {
		return false
	}

	m := make(map[int]int, 0)

	for i, n := range nums {
		if m[n] != 0 && i-(m[n]-1) <= k {
			return false
		}

		m[n] = i + 1
	}

	return false
}

func intersect(nums1 []int, nums2 []int) []int {
	res := []int{}
	m1 := getInts(nums1)
	m2 := getInts(nums2)

	if len(m1) > len(m2) {
		m1, m2 = m2, m1
	}
	// m1 是较短的字典，会快一些
	for n := range m1 {
		m1[n] = min(m1[n], m2[n])
	}

	for n, size := range m1 {
		for i := 0; i < size; i++ {
			res = append(res, n)
		}
	}

	return res

}

// 清理重复的值，也便于查询

func getInts(a []int) map[int]int {

	res := make(map[int]int, len(a))

	for i := range a {
		res[a[i]]++
	}
	return res
}

func min(a, b int) int {

	if a < b {
		return a
	}
	return b
}

// 方法二

func intersect1(nums1 []int, nums2 []int) []int {

	m := make(map[int]int, len(nums1))

	for _, j := range nums1 {
		m[j]++
	}

	k := 0
	// 在nums2 存在
	for _, j := range nums2 {
		if m[j] > 0 {
			m[j]--
			nums2[k] = j
			k++
		}
	}

	return nums2[0:k]

}

func in(nums []int, nums2 []int) []int {
	m := make(map[int]int, len(nums))

	for _, j := range nums {
		m[j]++
	}

	k := 0

	for _, j := range nums2 {
		if m[j] > 0 {
			m[j]--
			nums2[k] = j
			k++
		}
	}

	return nums2[0:k]
}
