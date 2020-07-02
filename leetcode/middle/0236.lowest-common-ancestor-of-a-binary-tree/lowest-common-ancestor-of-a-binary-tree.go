package _236_lowest_common_ancestor_of_a_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil ||
		root == p ||
		root == q {
		return root
	}

	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)

	// l=nil 意味着， p 和 q 都 **不在** root.Left 中
	// r=nil 同理。
	// 所以，根据题意， l 和 r 不可能同时为 nil
	if l != nil && r != nil {
		// 此时 p 和 q 分别在 root.Left 和 root.Right 中
		return root
	}
	if l == nil {
		// 此时 p 和 q 在 root.Right 中
		return r
	}
	// 此时 p 和 q 在 root.Left 中
	return l
}
