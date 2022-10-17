package _055_jump_game

func canJump(nums []int) bool {
	l := len(nums)
	for i := l - 1; i > 0; i-- {
		if nums[i]+i >= l {
			l = i
		}
	}
	return l <= 0
}

func canJump2(nums []int) bool {
	cover := nums[0]

	for i := 0; i < cover; i++ {
		if cover > nums[i]+i {
			return true
		}

		cover = max(cover, nums[i]+i)
	}

	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
