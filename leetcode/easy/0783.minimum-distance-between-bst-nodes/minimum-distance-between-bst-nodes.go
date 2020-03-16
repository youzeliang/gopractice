package _783_minimum_distance_between_bst_nodes

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var nodeVal []int

func minDiffInBST(root *TreeNode) int {
	nodeVal = make([]int, 0)
	minDiffInTool(root)
	min := 100

	for i := 1; i < len(nodeVal); i++ {
		tmp := nodeVal[i] - nodeVal[i-1]
		min = Min(min, tmp)
	}
	return min
}

func minDiffInTool(node *TreeNode) {
	if node == nil {
		return
	}
	if node.Left != nil {
		minDiffInTool(node.Left)
	}

	nodeVal = append(nodeVal, node.Val)

	if node.Right != nil {
		minDiffInTool(node.Right)
	}

}

func Min(a, b int) int {
	if a > b {
		return b
	}

	return a
}
