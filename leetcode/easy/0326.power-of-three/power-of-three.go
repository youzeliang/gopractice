package _326_power_of_three

/**
https://leetcode.com/problems/power-of-three/


给定一个整数，写一个函数来判断它是否是 3 的幂次方。


解题思路

可以不断除以 3 模 3


*/
func isPowerOfThree(n int) bool {
	if n < 1 {
		return false
	}

	for n%3 == 0 {
		n /= 3
	}

	return n == 1

}
