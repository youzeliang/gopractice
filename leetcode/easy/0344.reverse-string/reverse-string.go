package _344_reverse_string

// https://leetcode-cn.com/problems/reverse-string/
// 双指针交换

func reverseString(s []byte) {
	n := len(s)
	mid := n / 2
	for i := 0; i < mid; i++ {
		s[i], s[n-i-1] = s[n-i-1], s[i]
	}
}
