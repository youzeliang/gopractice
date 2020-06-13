package _376_wiggle_subsequence

// 使用状态机，当状态转变时，计数器 +1 ，在最后返回时再 +1，因为之前是状态转变的次数，比长度小1
func wiggleMaxLength(nums []int) int {
	count := len(nums)
	if count < 2 {
		return count
	}
	var ret int
	// 状态 0 初始，1 上升，2 下降
	var status int
	for i := 0; i < count-1; i++ {
		// 上升
		if nums[i] < nums[i+1] {
			if status == 2 || status == 0 {
				ret++
			}
			status = 1
		}
		// 下降
		if nums[i] > nums[i+1] {
			if status == 1 || status == 0 {
				ret++
			}
			status = 2
		}
	}
	ret++
	return ret
}