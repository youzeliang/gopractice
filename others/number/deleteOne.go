package number

// 删除k个数字后的最小值
// 给出一个整数，从该整数中去掉k个数字，要求剩下的数字形成的新整数尽可能的小

// 解题思路
// 把原整数的所有数字从左到右都进行比较，如果发现某一位的数字大于它右边的数字，那么在删除该数字后，必然会使该位数的值降低，因为右面比它小的数字顶替了它的位置

type Stack []byte

func (s *Stack) Push(v byte) {
	*s = append((*s), v)
}

func (s *Stack) Pop() byte {
	l := len(*s)
	v := (*s)[l-1]
	*s = (*s)[:l-1]
	return v
}

func (s *Stack) Top() byte {
	v := (*s)[len(*s)-1]
	return v
}

func (s *Stack) Len() int {
	return len(*s)
}

func removeKdigits(num string, k int) string {
	s := &Stack{}
	s.Push('0')

	for i := 0; i < len(num); i++ {
		for k > 0 && num[i] < s.Top() {
			s.Pop()
			k--
		}
		s.Push(num[i])
	}
	for k > 0 {
		s.Pop()
		k--
	}

	// Trim prefix '0'
	var i int
	for i < s.Len()-1 && (*s)[i] == '0' {
		i++
	}
	*s = (*s)[i:]

	return string(*s)
}
