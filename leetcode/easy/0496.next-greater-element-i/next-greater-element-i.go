package _496_next_greater_element_i

func nextGreaterElement(findNums []int, nums []int) []int {
	index := make(map[int]int)
	for i, n := range nums {
		index[n] = i
	}

	res := make([]int, len(findNums))
	for i, n := range findNums {
		res[i] = -1
		for j := index[n] + 1; j < len(nums); j++ {
			if n < nums[j] {
				res[i] = nums[j]
				break
			}
		}
	}

	return res
}
