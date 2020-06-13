package _357_count_numbers_with_unique_digits

//F(n) = 9P(n-1,9) + 9P(n-2,9) + ... + 9*P(1,9) + 10
//F(0) = 1

func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	sum := 0
	factor := 1
	for i := 1; i < n; i++ {
		factor *= (10 - i)
		sum += 9 * factor
	}
	return sum + 10
}