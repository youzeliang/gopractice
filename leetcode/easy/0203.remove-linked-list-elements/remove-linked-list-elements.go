package _203_remove_linked_list_elements

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	headPre := ListNode{Next: head}

	temp := &headPre
	for temp.Next != nil {
		if temp.Next.Val == val {
			// 删除符合条件的节点
			temp.Next = temp.Next.Next
		} else {
			temp = temp.Next
		}
	}

	return headPre.Next
}
