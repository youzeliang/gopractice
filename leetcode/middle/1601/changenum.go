package _601


func swapNumbers(numbers []int) []int {
	if len(numbers) != 2 {
		return numbers
	}

	numbers[0] = numbers[0] ^ numbers[1]
	numbers[1] = numbers[0] ^ numbers[1]
	numbers[0] = numbers[0] ^ numbers[1]

	return numbers

}
