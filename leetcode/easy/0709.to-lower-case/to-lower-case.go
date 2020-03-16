package _709_to_lower_case

import "bytes"

// https://leetcode-cn.com/problems/to-lower-case/solution/
// 大写字母比小写字母小32，把范围为(64,91)的元素加32得到它的小写。

func toLowerCase(str string) string {
	res := make([]byte, 0)
	for _, j := range str {
		if j < 91 && j > 64 {
			res = append(res, byte(j))
		}
	}

	return string(res)
}

func toLowerCase1(str string) string {
	var buf bytes.Buffer
	for _, c := range str {
		if c >= 65 && c <= 90 {
			buf.WriteRune(c + rune(32))
		} else {
			buf.WriteRune(c)
		}
	}
	return buf.String()
}
