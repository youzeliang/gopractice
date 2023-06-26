package main

import "fmt"

//给定一个二叉树，返回它的中序 遍历。
//
// 示例:
//
// 输入: [1,null,2,3]
//   1
//    \
//     2
//    /
//   3
//
//输出: [1,3,2]
//
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
// Related Topics 栈 树 哈希表

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归法
func inorderTraversal1(root *TreeNode) []int {

	return inorderRecursive(root)

}

func inorderRecursive(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	rest := append(inorderRecursive(root.Left), root.Val)
	rest = append(rest, inorderRecursive(root.Right)...)
	return rest
}

// 2. 迭代

//问题是典型 **深度优先搜索（DFS）** 模式，深度优先搜索改成 **迭代循环** 写法通常需借助 **栈** 来手动控制。

func inorderTraversal(root *TreeNode) []int {
	return inorderIterate(root)
}

func inorderIterate(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var rest []int
	stack := Stack([]*TreeNode{root})
	for len(stack) > 0 {
		cur := stack.Pop()
		if cur != nil {
			if cur.Right != nil {
				stack.Push(cur.Right)
			}
			stack.Push(cur)
			stack.Push(nil) // 待读取标记
			if cur.Left != nil {
				stack.Push(cur.Left)
			}
		} else {
			rest = append(rest, stack.Pop().Val)
		}
	}

	return rest
}

type Stack []*TreeNode

func (s *Stack) Push(node *TreeNode) {
	*s = append(*s, node)
}

func (s *Stack) Pop() *TreeNode {
	n := []*TreeNode(*s)[len(*s)-1]
	*s = []*TreeNode(*s)[:len(*s)-1]
	return n
}

func init() {
	fmt.Println("for some test11111")
}
