package _007_reverse_integer

import "math"

// https://leetcode-cn.com/problems/reverse-integer/

func reverse(x int) int {

	flag := 1

	// 处理负数
	if x < 0 {
		flag = -1
		x = flag * x
	}

	res := 0
	for x > 0 {
		// 取出x的末尾

		temp := x % 10
		// 放入 res 的开头

		res = res*10 + temp
		// x 去除末尾

		x = x / 10
	}

	// 还原 x 的符号到 res

	res = flag * res

	// 处理 res 的溢出问题
	if res > math.MaxInt32 || res < math.MinInt32 {
		res = 0
	}
	return res
}
