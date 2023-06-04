package _089

func duplicateZeros(arr []int) {
	dp := make([]int, len(arr))
	dp[0] = 0
	for i := 0; i < len(arr)-1; i++ {
		flag := 0
		if arr[i] == 0 {
			flag = 1
		}
		dp[i+1] = dp[i] + flag
	}
	for i := len(arr) - 1; i > 0; i-- {
		if (i + dp[i]) > (len(arr) - 1) {
			arr[i] = 0
			continue
		}
		if dp[i] > 0 {
			arr[i+dp[i]] = arr[i]
			arr[i] = 0
		}
	}
}
