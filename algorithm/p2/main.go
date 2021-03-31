package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func d(root *TreeNode, val int) *TreeNode {


	if root == nil {
		return nil
	}

	root.Left = d(root.Left, val)
	root.Right = d(root.Right, val)

	if root.Val == val && root.Left == nil && root.Right == nil {
		return nil
	}

	return root
}
