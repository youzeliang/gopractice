package main

func removeDuplicates(nums []int) int {
	n, p, s := len(nums), 1, 1
	for i := 1; i < n; i++ {
		if nums[i] > nums[p-1] {
			s = 1
		} else if s == 1 && nums[i] == nums[p-1] {
			s = 0
		} else {
			continue
		}
		nums[p] = nums[i]
		p++
	}
	return p
}
