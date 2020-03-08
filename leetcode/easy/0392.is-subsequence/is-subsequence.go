package _392_is_subsequence

// https://leetcode-cn.com/problems/is-subsequence/
// 双指针遍历

func isSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}

	t1, t2 := 0, 0
	for t1 < len(s) && t2 < len(t) {
		if s[t1] == t[t2] {
			t1++
		}

		t2++
	}

	return t1 == len(s)
}
