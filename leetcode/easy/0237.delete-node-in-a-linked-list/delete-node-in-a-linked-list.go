package _237_delete_node_in_a_linked_list

// ListNode is pre-defined...

type ListNode struct {
	Next *ListNode
	Val  int
}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
