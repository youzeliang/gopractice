package _611

func divingBoard(shorter int, longer int, k int) []int {
	if k == 0 {
		return []int{}
	}

	if shorter == longer {
		return []int{k}
	}

	var res []int

	for i := 0; i <= k; i++ {
		a := 0
		a = shorter*(k-i) + i*longer
		res = append(res, a)
	}

	return res
}
