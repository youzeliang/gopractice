package _find_

// https://leetcode-cn.com/problems/find-n-unique-integers-sum-up-to-zero/
func sumZero(n int) []int {
	pair := n / 2
	ans := make([]int, n)
	for i := 0; i < pair; i++ {
		ans[i] = i + 1
		ans[pair+i] = -(i + 1)
	}

	return ans
}
