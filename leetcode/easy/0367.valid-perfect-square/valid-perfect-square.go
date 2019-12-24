package _367_valid_perfect_square

func isPerfectSquare(num int) bool {
	l := num / 2
	if num <= 1 {
		return true
	}
	for i := 0; i <= l; i++ {
		if i*i == num {
			return true
		}
	}

	return false
}

// 二分法

func isPerfectSquareOne(num int) bool {
	if num <= 1 {
		return false

	}

	left, right := 1, num
	for left < right {
		mid := left + (right-left)/2
		if mid*mid > num {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return right*right == num

}

// 牛顿迭代法

func isPerfectSquareTwo(num int) bool {
	if num <= 0 {
		return false
	}

	r := num
	for r*r > num {
		r = (r + num/r) / 2
	}

	return r*r == num
}
