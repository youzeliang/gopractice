package _485_max_consecutive_ones

func findMaxConsecutiveOnes(nums []int) int {
	res, max := 0, 0

	for _, j := range nums {
		if j == 1 {
			res++
		}

		if max < res {
			max = res
		}

		if j == 0 {
			res = 0
		}

	}

	return max
}
