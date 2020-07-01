package _647_palindromic_substrings

// 回文子串个问题
// 第一时间要想到中间扩展法，这种解法应该算是最优解法
// 遍历一遍字符串，以当前字符为中间，向左右两边扩展扫描
// 同时也要考虑以当前字符和下一个字符组合为中心向左右两边扩展扫描

func extend(s string, l, r int) int {
	res := 0
	for l >= 0 && r < len(s) && s[l] == s[r] {
		res++
		r++
		l--
	}
	return res
}

func countSubstrings(s string) int {
	if len(s) == 0 {
		return 0
	}

	res := 0
	for i := 0; i < len(s); i++ {
		res += extend(s, i, i)
		res += extend(s, i, i+1)
	}
	return res
}
