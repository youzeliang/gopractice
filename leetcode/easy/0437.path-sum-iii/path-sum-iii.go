package _437_path_sum_iii

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	cnt := 0

	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}

		sum -= node.Val
		if sum == 0 {
			cnt++
		}

		dfs(node.Left, sum)
		dfs(node.Right, sum)
	}

	dfs(root, sum)

	return cnt +
		pathSum(root.Left, sum) +
		pathSum(root.Right, sum)
}
