package _617

func maxSubArray(nums []int) int {
	currentRes := 0
	maxRes := nums[0]

	for _, j := range nums {
		if currentRes < 0 {
			currentRes = j
		} else {
			currentRes += j
		}

		if currentRes > maxRes {
			maxRes = currentRes
		}
	}

	return maxRes
}
