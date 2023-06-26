package _392_is_subsequence

// https://leetcode-cn.com/problems/is-subsequence/
// 双指针遍历

func isSubsequence(s string, t string) bool {

	n, m := len(s), len(t)

	i, j := 0, 0

	for i < n && j < m {
		if s[i] == t[j] {
			i++
		}
		j++
	}

	return i == n

}
