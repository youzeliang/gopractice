package problem0009

import (
	"strconv"
)

// 转化为字符串来求
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	s := strconv.Itoa(x)

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}

	return true
}

//  利用不断求余
func isPalindrome1(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	r, o := 0, x
	for x != 0 {
		val := x % 10
		r = r*10 + val
		x = x / 10
	}
	return o == r
}
