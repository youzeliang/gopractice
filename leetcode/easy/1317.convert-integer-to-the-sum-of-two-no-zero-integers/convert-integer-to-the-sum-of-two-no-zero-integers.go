package _317_convert_integer_to_the_sum_of_two_no_zero_integers

import (
	"strconv"
	"strings"
)

// https://leetcode-cn.com/problems/convert-integer-to-the-sum-of-two-no-zero-integers/
func getNoZeroIntegers(n int) []int {
	var slow int = 1
	var fast int = n - 1
	var res = make([]int, 0)
	for slow <= fast {
		var strSlow = strconv.Itoa(slow)
		var strFast = strconv.Itoa(fast)
		if !strings.Contains(strSlow, "0") && !strings.Contains(strFast, "0") {
			res = append(res, slow, fast)
			break
		}
		slow++
		fast--
	}
	return res
}
