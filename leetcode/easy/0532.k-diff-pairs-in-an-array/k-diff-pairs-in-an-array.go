package _532_k_diff_pairs_in_an_array

func findPairs(nums []int, k int) int {
	if k < 0 {
		return 0
	}

	numsMap := make(map[int]bool)
	diffMap := make(map[int]bool)

	for _, j := range nums {
		if numsMap[j-k] {
			diffMap[j-k] = true
		}

		if numsMap[j+k] {
			diffMap[j] = true
		}

		numsMap[j] = true
	}

	return len(diffMap)
}
