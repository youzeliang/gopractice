package _710

// https://leetcode-cn.com/problems/find-majority-element-lcci/

func majorityElement(nums []int) int {
	numLens := len(nums)
	res := make(map[int]int, numLens)
	for _, v := range nums {
		res[v]++
	}

	for k, v := range res {
		if v > numLens/2 {
			return k
		}
	}

	return -1
}
