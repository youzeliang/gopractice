package _693_binary_number_with_alternating_bits

func hasAlternatingBits(n int) bool {
	// 解法1,只需要判断是否有连续的"11"或者"00"即可
	// nStr := fmt.Sprintf("%b", n)
	// return !strings.Contains(nStr, "11") && !strings.Contains(nStr, "00")

	// 解法2
	// 将n右移一位异或n，检查结果是否全为1
	tmp := n ^ (n >> 1)
	return tmp&(tmp+1) == 0
}
