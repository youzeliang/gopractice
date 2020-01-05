package _633_sum_of_square_numbers

import "math"

func judgeSquareSum(c int) bool {
	a := getIntSqrt(c)

	for a >= 0 {
		if isSquare(c - a*a) {
			return true
		}

		a--
	}

	return false
}

func isSquare(c int) bool {
	x := getIntSqrt(c)
	return x*x == c
}

func getIntSqrt(c int) int {
	return int(math.Sqrt(float64(c)))
}
