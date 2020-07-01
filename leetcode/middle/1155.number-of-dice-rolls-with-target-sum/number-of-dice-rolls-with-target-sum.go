package _155_number_of_dice_rolls_with_target_sum

func numRollsToTarget(d int, f int, target int) int {
	mod := int((1e9) + 7)
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 0; i < d; i++ {
		ndp := make([]int, target+1)
		for j := 1; j <= f; j++ {
			for k := 0; k <= target-j; k++ {
				ndp[k+j] = (ndp[k+j] + dp[k]) % mod
			}
		}
		//fmt.Println(dp)
		dp = ndp
	}
	return dp[target]
}
