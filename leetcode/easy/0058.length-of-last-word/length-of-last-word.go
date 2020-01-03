package _058_length_of_last_word

func lengthOfLastWord(s string) int {
	res := 0
	l := len(s)
	if l == 0 {
		return 0
	}

	for i := l - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if res != 0 {
				return res
			}
			continue
		}

		res++

	}

	return res
}
