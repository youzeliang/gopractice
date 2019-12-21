package _221_split_a_string_in_balanced_strings

// https://leetcode-cn.com/problems/split-a-string-in-balanced-strings/
// 其实只需要统计'R'与'L'的数量，当两个字符的数量相等时就表示可以进行分割，这里不需要对已统计的清零，，因为每次划分时两个字符数量是相等的

func balancedStringSplit(s string) int {
	res, l, r := 0, 0, 0
	start := s[0]
	for i := range s {
		if s[i] == start {
			l++
		} else {
			r++
		}
		if l == r {
			res++
		}
	}

	return res
}
