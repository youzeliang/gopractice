package _102


// https://leetcode-cn.com/problems/check-permutation-lcci/

func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	var res byte  = 0


	for i := range s1{
		res ^= s1[i]
	}

	for i := range s2{
		res ^= s2[i]
	}

	return res == 0
}


