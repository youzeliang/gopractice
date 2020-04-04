package _58

func reverseLeftWords(s string, n int) string {
	tmpRune := []rune(s)
	if n > len(tmpRune) {
		return s
	}
	result := append(tmpRune[n:], tmpRune[:n]...)
	return string(result)
}
