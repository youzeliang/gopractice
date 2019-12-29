package _205_isomorphic_strings

import "bytes"

// 判断第一次出现的位置
func isIsomorphic(s string, t string) bool {
	ss := []byte(s)
	tt := []byte(t)

	if len(ss) == 0 && len(tt) == 0 {
		return true
	}

	if len(ss) != len(tt) {
		return false
	}

	for i := 0; i < len(ss); i++ {
		if bytes.IndexByte(ss, s[i]) != bytes.IndexByte(tt, tt[i]) {
			return false
		}
	}

	return true

}
