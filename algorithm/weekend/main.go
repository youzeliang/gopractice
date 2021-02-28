package main

import "sort"

func main() {
	tupleSameProduct([]int{2, 3, 4, 6})
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
