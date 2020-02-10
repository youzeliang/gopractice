package _003_longest_substring_without_repeating_characters

import "strings"

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
