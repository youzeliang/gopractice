package _066_plus_one

// append(a, b...)  // 合并切片 a 和 b，b... 意思为将切片 b 转换为参数列表
func plusOne(digits []int) []int {

	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			// 最末尾的一位，直接返回
			digits[i]++
			return digits
		} else {
			// 置为0
			digits[i] = 0
		}
	}

	return append([]int{1}, digits...)
}
