package _709_to_lower_case

import (
	"fmt"
	"testing"
)

func Test_a(t *testing.T) {

	str := "abc"

	strRune := []rune(str)
	fmt.Print(strRune)

	fmt.Println(string(strRune))
}
