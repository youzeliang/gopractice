package _788_rotated_digits

func rotatedDigits(N int) int {
	count := 0
	for i := 2; i <= N; i++ {
		if isValid(i) {
			count++
		}
	}

	return count
}

func isValid(n int) bool {
	var isValid bool
	for n > 0 {
		switch n % 10 {
		case 2, 5, 6, 9:
			isValid = true
		case 3, 4, 7:
			return false
		}
		n /= 10
	}

	return isValid
}
