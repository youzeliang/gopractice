package _60

func singleNumber(arr []int) []int {
	var result []int
	if arr == nil {
		return result
	}
	xor := 0
	// 计算数组中所有数字异或的结果
	for _, v := range arr {
		xor ^= v
	}
	position := 0
	for i := xor; i&1 == 0; i >>= 1 {
		position++
	}
	// 异或的结果与所有第position位为1的数字异或
	// 结果一定是出现一次的两个数中的一个
	tempXor := xor
	for _, v := range arr {
		if (uint(v)>>uint(position))&1 == 1 {
			xor ^= v
		}
	}
	// 得到另外一个出现一次的数字
	return []int{xor, xor ^ tempXor}
}
