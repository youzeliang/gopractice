package problem0001

import (
	"fmt"
	"strconv"
	"testing"
)

type para struct {
	one []int
	two int
}

type ans struct {
	one []int
}

type question struct {
	p para
	a ans
}

func Test_OK(t *testing.T) {
	a := 4141
	for a != 0 {

		a /= len(strconv.Itoa(a))*10

		fmt.Println(a)
		break

	}

	//a %= 10
	//a /=10
	//println(a)
}
