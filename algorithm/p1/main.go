package main

import (
	"fmt"
	"strings"
)

func main() {



	//fmt.Println(runtime.NumCPU())
	//
	//
	//numJewelsInStones("Add","Addccc")


	//intersect(nil)

	arr := []int{1, 2, 3}
	for _, v := range arr {
		arr = append(arr, v)
	}
	fmt.Println(arr)

}

func consecutiveNumbersSum(N int) int {

	var res  = 1
	var sum = 1
	for i := 2;sum<=n;i++{
		sum += i
		if (N-sum) % i == 0 {
			res++
		}
	}

	return res
}

func intersect(nums []int  )  {

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

