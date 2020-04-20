package _203_delete_middle_node_lcci

type ListNode struct {
	Val  int
	Next *ListNode
}


func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
