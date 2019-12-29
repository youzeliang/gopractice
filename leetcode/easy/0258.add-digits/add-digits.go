package _258_add_digits

// 根数字等于原数除以9的余数，这个过程叫做弃9法，这里考虑到
func addDigits(n int) int {
	return (n-1)%9 + 1
}
