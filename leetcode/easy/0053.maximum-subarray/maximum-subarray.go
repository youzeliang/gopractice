package _053_maximum_subarray

func maxSubArray(nums []int) int {
	sum, maxSum := -1<<31, -1<<31
	for _, n := range nums {
		// sum+n < n，那就还不如直接从 n 开始统计。
		sum = max(sum+n, n)
		maxSum = max(maxSum, sum)
	}
	return maxSum

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 解法2

func maxSubArray2(nums []int) int {
	curMaxSum := 0       // 当前最大连续子序列和
	allMaxSum := nums[0] // 全局最大连续子序列和
	for _, x := range nums {
		if curMaxSum < 0 { // 当前最大和如果小于0，直接用当前元素值覆盖这个负值
			curMaxSum = x
		} else { // 如果当前最大和>=0，继续累加
			curMaxSum += x
		}
		if curMaxSum > allMaxSum {
			allMaxSum = curMaxSum
		}
	}
	return allMaxSum
}
