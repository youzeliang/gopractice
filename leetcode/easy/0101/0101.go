package _101

// https://leetcode-cn.com/problems/is-unique-lcci/
func isUnique(astr string) bool {
	var mark uint32 = 0
	for _, ch := range astr {
		move_bit := ch - 'a'
		if mark&(1<<move_bit) != 0 {
			return false
		} else {
			mark |= (1 << move_bit)
		}
	}
	return true
}
