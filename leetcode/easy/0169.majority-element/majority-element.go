package _169_majority_element

import (
	"sort"
)

// https://leetcode-cn.com/problems/majority-element/

func majorityElement(nums []int) int {
	sort.Ints(nums)
	i := len(nums)/2
	return nums[i]
}
