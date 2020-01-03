package _342_power_of_four

func isPowerOfFour(num int) bool {
	return num > 0 && num&(num-1) == 0 && 0xaaaaaaaa&num == 0
}
