package _338_counting_bits

// 动态规划 dp[i]表示i中1的位数  1. dp[i]=dp[i/2],偶数  2. dp【i】=dp[i-1]+1  奇数
func countBits(num int) []int {
	dp := make([]int, num+1)
	dp[0] = 0
	for i := 0; i < num+1; i++ {
		// 偶数
		if i%2 == 0 {
			dp[i] = dp[i/2]
		} else {
			dp[i] = dp[i-1] + 1
		}
	}
	return dp
}