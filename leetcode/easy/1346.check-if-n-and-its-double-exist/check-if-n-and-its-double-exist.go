package _346_check_if_n_and_its_double_exist

// https://leetcode-cn.com/problems/check-if-n-and-its-double-exist/
func checkIfExist(arr []int) bool {
	tmp := map[int]int{}
	for _, v := range arr {
		if tmp[v*2] == 1 || (v&1 == 0 && tmp[v/2] == 1) {
			return true
		}
		tmp[v] = 1
	}
	return false
}
