package _728_self_dividing_numbers

func selfDividingNumbers(left int, right int) []int {
	var res []int
	for left := left; left <= right; left++ {
		if jump(left) {
			res = append(res, left)
		}
	}

	return res
}

func jump(n int) bool {
	t := n
	for t > 0 {
		d := t % 10
		t /= 10
		if d == 0 || n%d != 0 {
			return false
		}
	}

	return true

}
