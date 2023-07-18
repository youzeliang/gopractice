package _055_jump_game

func calJump(nums []int) bool {

	l := len(nums)

	for i := l - 1; i >= 0; i-- {
		if nums[i]+i >= l {
			l = i
		}
	}

	return l <= 0
}

func jump(num []int) bool {

	//cover := num[0]

	//for i := 0; i <=cover;i++ {
	//	if cover >= len(num) -1 {
	//		return true
	//	}
	//
	//	cover = max(cover i + num[i])
	//}

	return false
}

func max(a, b int) int {
	if a > b {

	}

	return 0
}
