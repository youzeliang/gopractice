package _402

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return tool(nums)
}

func tool(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	} else {
		root := new(TreeNode)
		root.Val = nums[len(nums)/2]
		root.Left = tool(nums[0:(len(nums) / 2)])
		root.Right = tool(nums[(len(nums)/2)+1:])
		return root

	}

}
