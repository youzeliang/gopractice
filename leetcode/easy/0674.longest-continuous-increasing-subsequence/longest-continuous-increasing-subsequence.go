package _674_longest_continuous_increasing_subsequence

func findLengthOfLCIS(nums []int) int {
	res, max := 1, 1
	l := len(nums)
	if l == 0 {
		return 0
	}

	for i := 1; i < l; i++ {
		if nums[i] > nums[i-1] {
			res++
		} else {
			res = 1
		}
		if res > max {
			max = res
		}
	}
	return max
}
