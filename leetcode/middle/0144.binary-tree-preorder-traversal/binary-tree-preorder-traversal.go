package _144_binary_tree_preorder_traversal

var res []int

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
func preorderTraversal(root *TreeNode) []int {
	res = []int{}
	dfs(root)
	return res
}

func dfs(root *TreeNode) {
	if root != nil {
		res = append(res, root.Val)
		dfs(root.Left)
		dfs(root.Right)
	}
}

//迭代 （stack）
func preorderTraversal1(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode

	for 0 < len(stack) || root != nil { //root != nil 只为了第一次root判断，必须放最后
		for root != nil {
			res = append(res, root.Val)       //前序输出
			stack = append(stack, root.Right) //右节点 入栈
			root = root.Left                  //移至最左
		}
		index := len(stack) - 1 //栈顶
		root = stack[index]     //出栈
		stack = stack[:index]
	}
	return res
}

// Morris (树转链表)
func preorderTraversal2(root *TreeNode) []int {
	var max *TreeNode
	var res []int
	for root != nil {
		if root.Left == nil {
			res = append(res, root.Val)
			root = root.Right
		} else {
			max = root.Left
			for max.Right != nil {
				max = max.Right
			}

			root.Right, max.Right = root.Left, root.Right
			root.Left = nil
		}
	}
	return res
}

// Morris (保持树结构)
func preorderTraversal3(root *TreeNode) []int {
	var max *TreeNode
	var res []int
	for root != nil {
		if root.Left == nil {
			res = append(res, root.Val)
			root = root.Right
		} else {
			max = root.Left
			for max.Right != nil && max.Right != root {
				max = max.Right
			}

			if max.Right == nil {
				res = append(res, root.Val)
				max.Right = root.Right
				root = root.Left
			} else {
				root = root.Right
				max.Right = nil
			}
		}
	}
	return res
}
