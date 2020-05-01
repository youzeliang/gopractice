package _016_3sum_closest

import (
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	res := nums[0] + nums[1] + nums[2]

	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		left, right := i+1, len(nums)-1
		for left < right {
			tem := nums[i] + nums[left] + nums[right]
			if tem > target {
				right--
			} else if tem < target {
				left++
			} else {
				return target
			}
			if distance(tem, target) < distance(res, target) {
				res = tem
			}

		}
	}

	return res

}

func distance(a, b int) int {
	if a < b {
		return b - a
	}

	return a - b
}
