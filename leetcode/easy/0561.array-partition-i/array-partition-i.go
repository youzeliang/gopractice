package _561_array_partition_i

import "sort"

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	res := 0
	l := len(nums)
	for i := 0; i < l; i += 2 {
		res += nums[i]
	}

	return res
}