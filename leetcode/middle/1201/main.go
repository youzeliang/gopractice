package main

func nthUglyNumber(n int, a int, b int, c int) int {
	ab, ac, bc := lcm(a, b), lcm(a, c), lcm(b, c)
	abc := lcm(ab, c)
	left, right := 0, n*min(a, b, c)
	for left <= right {
		mid := (left + right) / 2
		num := mid/a + mid/b + mid/c - mid/ab - mid/ac - mid/bc + mid/abc
		if num == n {
			if mid%a == 0 || mid%b == 0 || mid%c == 0 {
				return mid
			} else {
				right = mid - 1
			}
		} else if num > n {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

// 求最小数
func min(a int, args ...int) (result int) {
	result = a
	for _, v := range args {
		if v < result {
			result = v
		}
	}
	return result
}

// 求最小公倍数
func lcm(x, y int) int {
	return x * y / gcd(x, y)
}

// 求最大公约数
func gcd(x, y int) int {
	tmp := x % y
	if tmp > 0 {
		return gcd(y, tmp)
	}
	return y
}
