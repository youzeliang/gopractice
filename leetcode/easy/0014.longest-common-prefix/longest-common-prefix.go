package _014_longest_common_prefix

func longestCommonPrefix(strs []string) string {
	short := shortest(strs)

	for i, j := range short {
		for m := 0; m < len(strs); m++ {
			if strs[m][i] != byte(j) {
				return strs[m][:i]
			}
		}
	}

	return short
}

// 求出最短字符串
func shortest(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	res := strs[0]
	for _, j := range strs {
		if len(j) < len(res) {
			res = j
		}
	}

	return res
}

<<<<<<< HEAD




=======
func daysBetweenDates(date1 string, date2 string) int {

	a, _ := time.Parse(date1, date1)
	println(a)

}
>>>>>>> a5e0a62349c7c711a13bf0e3c24c2a7432635906
