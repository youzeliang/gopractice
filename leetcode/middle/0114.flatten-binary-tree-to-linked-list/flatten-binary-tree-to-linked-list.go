package _114_flatten_binary_tree_to_linked_list

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 题目要求原地展开，故不能使用数组类变量存储全部节点值再重构树。将左右子树分别递归展开，将原左子树变为节点的右子树，再将原右子树变为当前右子树最右节点的右子树。

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	r := root.Right
	root.Right, root.Left = root.Left, nil
	for root.Right != nil {
		root = root.Right
	}
	root.Right = r

}
