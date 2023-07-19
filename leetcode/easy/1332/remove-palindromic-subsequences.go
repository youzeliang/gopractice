package _332

func removePalindromeSub(s string) int {
	if len(s) == 0 {
		return 0
	}

	if judge(s) {
		return 1
	}

	return 2
}

// 判断回文数
func judge(str string) bool {

	s := []rune(str)

	for key := range s {
		if key == (len(s) / 2) {
			break
		}

		last := len(s) - key - 1

		if s[key] != s[last] {
			return false
		}
	}

	return true
}
