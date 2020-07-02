package _222_count_complete_tree_nodes

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	count := 0
	traverse(root, &count)
	return count
}

func traverse(n *TreeNode, count *int) {
	if n == nil {
		return
	}

	*count++

	traverse(n.Left, count)
	traverse(n.Right, count)
}
