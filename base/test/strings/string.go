package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(isUnique("fsfsd"))

	var a = [][]int{{0, 0}, {1, 2}, {2, 4}, {3, 6}, {4, 8}}
	diagonalSum(a)

}

func Split(s, sep string) (result []string) {
	i := strings.Index(s, sep)

	for i > -1 {
		result = append(result, s[:i])
		s = s[i+1:]
		i = strings.Index(s, sep)
	}
	result = append(result, s)
	return
}

func isUnique(astr string) bool {
	var mark uint32 = 0
	for _, ch := range astr {
		moveBit := ch - 'a'
		if mark&(1<<moveBit) != 0 {
			return false
		} else {
			mark |= 1 << moveBit
		}
	}
	return true
}

func diagonalSum(mat [][]int) int {
	ans := 0
	for i := 0; i < len(mat); i++ {
		ans += mat[i][i]
		if i != len(mat)-1-i {
			ans += mat[i][len(mat)-1-i]
		}
	}
	return ans

}
