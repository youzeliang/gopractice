package _263_ugly_number

func isUgly(num int) bool {
	if num <= 0 {
		return false
	}

	if num <= 6 {
		return true
	}

	if num%2 == 0 {
		return isUgly(num / 2)
	}

	if num%3 == 0 {
		return isUgly(num / 3)
	}

	if num%5 == 0 {
		return isUgly(num / 5)
	}

	return false
}
