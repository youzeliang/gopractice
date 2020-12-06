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

func removeElement2(nums []int, val int) int {
	for i := 0; i < len(nums); {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}

	return len(nums)
}