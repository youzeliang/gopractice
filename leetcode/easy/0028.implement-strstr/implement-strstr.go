package _028_implement_strstr

func strStr(haystack string, needle string) int {

	h := len(haystack)
	n := len(needle)

	if n == 0 {
		return 0
	}

	if h < n {
		return -1
	}

	for i := 0; i <= h-n; i++ {
		if haystack[i:i+n] == needle {
			return i
		}
	}

	return -1

}
