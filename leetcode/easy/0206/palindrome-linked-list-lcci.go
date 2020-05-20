package _206

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	var res []*ListNode
	for head != nil {
		res = append(res, head)
		head = head.Next
	}
	left := 0
	right := len(res) - 1
	for left < right {
		if res[left].Val != res[right].Val {
			return false
		}
		left++
		right--
	}
	return true
}
