package _264_ugly_number_ii

func nthUglyNumber(n int) int {
	if n == 0 {
		return 1
	}
	dp := make([]int, n)
	dp[0] = 1
	idx2, idx3, idx5 := 0, 0, 0
	for i := 1; i < n; i++ {
		dp[i] = min(dp[idx2]*2, dp[idx3]*3, dp[idx5]*5)
		if dp[i] == dp[idx2]*2 {
			idx2++
		}
		if dp[i] == dp[idx3]*3 {
			idx3++
		}
		if dp[i] == dp[idx5]*5 {
			idx5++
		}
	}
	return dp[n-1]
}

func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}

	if b <= a && b <= c {
		return b
	}

	return c
}
