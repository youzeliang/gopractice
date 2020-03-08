package _803

import "math"

func findMagicIndex(nums []int) int {

	res := 0
	var min = math.MaxInt32
	for i := range nums {
		if i == nums[i] {
			res = i
		} else {
			continue
		}
		if res <= min {
			min = res
		}
	}

	if min == math.MaxInt32 {
		return -1
	}

	return min

}

// 更优的解法

func findMagicIndex2(nums []int) int {

	for i := 0; i < len(nums); i++ {
		if nums[i] == i {
			return i
		} else if nums[i] > i { // 如果nums[i] > i， 说明 i 至少为 nums[i] 时才可能有结果
			i = nums[i] - 1
		}
	}
	return -1

}
