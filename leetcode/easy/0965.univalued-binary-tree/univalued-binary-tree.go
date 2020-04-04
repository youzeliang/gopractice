package _965_univalued_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isTree(root, root.Val)

}

func isTree(root *TreeNode, val int) bool {
	if root == nil {
		return true
	}
	if root.Val != val {
		return false
	}
	return isTree(root.Left, val) && isTree(root.Right, val)
}
