package _5

func jump(nums []int) int {
	// 动态规划四步走：
	// 1、状态：f[i] 表示从起点到当前位置跳的最小次数
	// 2、推导：f[i] = min(f[j]+1),a[j])  j >=i 表示从j位置用一步跳到当前位置，这个j位置可能有多个，取最小的一个就行
	// 3、初始化：f[0] = 0
	// 4、结果：f[n-1]
	f := make([]int, len(nums))
	f[0] = 0
	for i := 1; i < len(nums); i++ {
		// f[i] 先取一个默认最大值i
		f[i] = i
		// 遍历之前结果取一个最小值+1
		for j := 0; j < i; j++ {
			if nums[j]+j >= i {
				f[i] = min(f[j]+1, f[i])
			}
		}
	}
	return f[len(nums)-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//贪心
//
//设每次所选的起跳位置为start，
//能到的最大距离end为i+nums[start]，
//我们为了能用最小的次数跳到终点，每次都想跳到最大的位置为max。
//跳的次数为count
//
//为了保证count最小，所以每次跳都要这一跳能跳的最远位置，解题思路应为贪心算法。从前向后遍历，
//每完成一次跳跃后更新count与end的值，下一次的i即从上一跳的最远处end_{i-1}(=start_i)endi−1(=starti)开始计算，
//遍历位置start_istarti到end_iendi得到能跳的最远处max，然后当遍历到end_iendi时，更新下一跳的起点start_{i+1}starti+1为max

func jump2(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}

	n := len(nums)
	count, max, end := 0, 0, 0
	for i := 0; i < n-1; i++ {

		// 记录每一个位置i所能跳到的最大位置max
		if i+nums[i] > max {
			max = i + nums[i]
		}
		// end是上一个所选的起跳位置能跳的最远处，当到达end时，更新下一跳的最大位置
		if i == end {
			end = max
			count++
		}
	}
	return count
}
