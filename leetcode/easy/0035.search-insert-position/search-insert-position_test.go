package _035_search_insert_position

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type question struct {
	para
	ans
}

// para 是参数
// one 代表第一个参数
type para struct {
	one []int
	two int
}

// ans 是答案
// one 代表第一个答案
type ans struct {
	one int
}

func Test_num35(t *testing.T) {

	ast := assert.New(t)

	qs := []question{
		{
			para{[]int{1, 3, 5, 6}, 5},
			ans{2},
		},

		{
			para{[]int{1, 3, 5, 6}, 2},
			ans{1},
		},

		{
			para{[]int{1, 3, 5, 6}, 7},
			ans{4},
		},

		{
			para{[]int{1, 3, 5, 6}, 0},
			ans{0},
		},

		{
			para{[]int{1, 3, 5, 6}, 6},
			ans{3},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, searchInsert(p.one, p.two), "输入:%v", p)
	}
}
