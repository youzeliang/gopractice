package _281

// https://leetcode-cn.com/problems/subtract-the-product-and-sum-of-digits-of-an-integer/
func subtractProductAndSum(n int) int {
	sum := 0
	product := 1
	for n > 0 {
		temp := n % 10
		sum += temp
		product *= temp
		n /= 10
	}

	return product - sum

}
