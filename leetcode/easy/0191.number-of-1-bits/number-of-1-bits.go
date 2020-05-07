package _191_number_of_1_bits

func HammingWeight(num uint32) int {
	count := 0
	for num != 0 {
		count++
		num &= num - 1
		// assume num = 0b1000
		// num-1 = 0b0111
		// num & (num-1) = 0
	}
	return count
}
