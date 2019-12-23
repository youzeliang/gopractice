package _013_roman_to_integer

func romanToInt(s string) int {
	res := 0

	l := len(s)
	for i := 0; i < l; i++ {
		if i+1 < l {
			if s[i] < s[i+1] {
				res -= letterToInt(string(s[i]))
			} else {
				res = letterToInt(string(s[i]))
			}
		}

	}
	return res

}

func letterToInt(s string) int {
	switch s {
	case "I":
		return 1
	case "V":
		return 5
	case "X":
		return 10
	case "L":
		return 50
	case "C":
		return 100
	case "D":
		return 500
	case "M":
		return 5000
	}

	return 0
}
