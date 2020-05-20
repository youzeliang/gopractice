package _367_linked_list_in_binary_tree

type ListNode struct {
	Val  int
	Next *ListNode
}
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubPath(head *ListNode, root *TreeNode) bool {
	if root == nil {
		return false
	}
	return isSub(head, root) || isSubPath(head, root.Left) || isSubPath(head, root.Right)
}

func isSub(head *ListNode, root *TreeNode) bool {
	if head == nil {
		return true
	}
	if root == nil || root.Val != head.Val {
		return false
	}
	return isSub(head.Next, root.Left) || isSub(head.Next, root.Right)
}
