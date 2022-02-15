package main

import (
	"fmt"
	"sort"
)

func main() {

	var a = []int{1,0,0,0,0,0,0,0, 2, 2, 8}

	var b = []int{6, 6}

	var k int
	for i := 0; i < len(a); i++ {
		if a[i] == 0 && len(b) > k  {
			a[i] = b[k]
			k++
		}
	}

	if k <= len(b) {
		a = append(a, b[k:]...)
	}

	fmt.Println(a)

}

func tupleSameProduct(nums []int) int {
	if len(nums) > 4 {
		nums = append(nums, nums[:len(nums)-1]...)
	}

	res := 0
	for i := 0; i <= len(nums)-4; i++ {
		if c(nums[i : i+4]) {
			res++
		}
	}

	return res * 8

}

func c(nums []int) bool {
	sort.Ints(nums)
	return nums[0]*nums[3] == nums[1]*nums[2]
}

// 输入：nums = [2,3,4,6,8,12]
