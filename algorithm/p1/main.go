package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	ss()
}

func m(nums []int) int {

	var count, major int

	for _, m := range nums {

		if count == 0 {
			major = m
		}

		if major == m {
			count++
		} else {
			count--
		}

	}

	return major

}

func ss() {

	s := [][]int{{1, 3},
		{2, 6},
		{15, 18},
		{8, 10},
	}

	fmt.Println(s)

	sort.Slice(s, func(i, j int) bool {
		return s[i][0] < s[j][0]
	})

	fmt.Println(s)

}

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 0 {
		return nil
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var merged [][]int
	merged = append(merged, intervals[0])

	for i := 1; i < len(intervals); i++ {
		m := merged[len(merged)-1]
		c := intervals[i]

		if c[0] > m[1] {
			merged = append(merged, c)
			continue
		}

		if c[1] > m[1] {
			m[1] = c[1]
		}
	}

	return merged
}

func canPartitionKSubsets(nums []int, k int) bool {
	size := len(nums)
	sum := 0
	max := nums[0]
	for _, n := range nums {
		sum += n
		if max < n {
			max = n
		}
	}

	if sum%k != 0 || sum/k < max {
		// 提前结束
		return false
	}

	// nums 降序排列可以加快 dfs 的速度
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	// isUsed[i] == true 表示 nums[i] 已经被归类到某个小组，无法再使用了
	isUsed := make([]bool, size)

	return dfs(nums, isUsed, k, 0, 0, sum/k)
}

func dfs(nums []int, isUsed []bool, k, start, sum, target int) bool {
	if k == 1 {
		// 已经找到了 k-1 组解
		// 剩下的自然就是第 k 组解
		return true
	}

	if sum == target {
		// 找到第 k 组的一种解
		// 开始搜索 k-1 组的解
		return dfs(nums, isUsed, k-1, 0, 0, target)
	}

	for i := start; i < len(nums); i++ {
		if !isUsed[i] && sum+nums[i] <= target {
			isUsed[i] = true
			// 试着 sum+nums[i]
			if dfs(nums, isUsed, k, i+1, sum+nums[i], target) {
				return true
			}
			isUsed[i] = false
		}
	}

	return false
}

func intersect(nums []int) {

	s1 := "hello"
	// 强制类型转换
	byteS1 := []byte(s1)

	fmt.Println(byteS1)
	byteS1[0] = 'H'
	fmt.Println(string(byteS1))
	s2 := "博客"
	runeS2 := []rune(s2)
	runeS2[0] = '狗'
	fmt.Println(string(runeS2))

}

func numJewelsInStones(J string, S string) int {
	isJewel := make(map[byte]bool, len(J))
	for i := range J {
		isJewel[J[i]] = true
	}

	res := 0
	for i := range S {
		if isJewel[S[i]] {
			res++
		}
	}

	return res
}

// 为经典的窗口滑动中的非定长算法

// 依次遍历字符串，当前位置字符串如果包含在窗口中，窗口中的开始位置更新为重复位置的下一位 最后得到最长非重复字符串的值

func lengthOfLongestSubstring(s string) int {
	// 定义游标尺寸大小,游标的左边位置
	window, start := 0, 0

	for i := 0; i < len(s); i++ {

		// 查看字符串是否在游标内
		isExist := strings.Index(string(s[start:i]), string(s[i]))

		// 如果不存在游标内部,游标长度重新计算并赋值
		if isExist == -1 {
			if i-start+1 > window {
				window = i - start + 1
			}
		} else {
			// 存在，游标开始位置更换为重复字符串位置的下一个位置
			start = start + 1 + isExist
		}
	}
	return window
}
