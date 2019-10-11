package _268_missing_number

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type question struct {
	para
	ans
}

type para struct {
	nums []int
}

type ans struct {
	one int
}

func Test_num0268(t *testing.T) {

	ast := assert.New(t)

	qs := []question{
		{
			para{[]int{0, 1, 2}},
			ans{3},
		},

		question{
			para{
				[]int{5, 4, 0, 1, 3},
			},
			ans{
				2,
			},
		},

		question{
			para{
				[]int{0, 1, 3},
			},
			ans{
				2,
			},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)
		ast.Equal(a.one, missingNumber(p.nums), "输入:%v", p)

	}

}
