package _027_remove_element

func removeElement(nums []int, val int) int {
	result := 0
	for _, j := range nums {
		if j != val {
			nums[result] = j
			result++
		}
	}

	return result
}
