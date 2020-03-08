package _106

// https://leetcode-cn.com/problems/compress-string-lcci/

import "strconv"

func compressString(S string) string {

	if len(S) <= 1 {
		return S
	}

	s := ""

	j := 1
	for i := 1; i < len(S); i++ {
		if S[i] == S[i-1] {
			j++
		} else {
			s = s + string(S[i-1]) + strconv.Itoa(j)
			j = 1
		}
	}
	s = s + string(S[(len(S)-1)]) + strconv.Itoa(j)

	if len(s) >= len(S) {
		return S
	}

	return s
}
