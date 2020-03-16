package _137_n_th_tribonacci_number

func tribonacci(n int) int {
	if n < 2 {
		return n
	}

	a, b, c := 0, 1, 1
	for n > 2 {
		a, b, c = b, c, a+b+c
		n--
	}

	return c
}

func tribonacci2(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	if n == 2 {
		return 1
	}

	dp := make([]int, n)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 1

	for i := 3; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
	}

	return dp[n-1] + dp[n-2] + dp[n-3]

}
