package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(isUnique("fsfsd"))
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
