package _389_find_the_difference

// https://leetcode-cn.com/problems/find-the-difference/
// 字符串转化为整数，然后相减

func findTheDifference(s string, t string) byte {
	sums := 0
	for _, i := range s {
		sums += int(i)
	}

	sumt := 0
	for _, i := range t {
		sumt += int(i)
	}

	return byte(sumt - sums)
}
