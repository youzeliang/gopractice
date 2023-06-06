package _111z

// 1，声明一个int型变量depth和整型数组变量ans
// 2，遍历字符串seq，当字符为(时深度递增，同时将深度值append进ans；否则字符就是)，将字符为)的深度值append进ans，同时深度值-1
// 3，返回最终结果ans

func maxDepthAfterSplit(seq string) []int {
	//1,use stack
	depth := 0
	var ans []int
	for _, c := range seq {
		if c == '(' {
			depth++
			ans = append(ans, depth%2)
		} else {
			ans = append(ans, depth%2)
			depth--
		}
	}
	return ans
}
