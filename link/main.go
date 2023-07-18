package link

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	res := &ListNode{
		Val: 1,
	}

	fmt.Println(res)

	if l1 == nil || l2 == nil {

	}

	return nil
}
