package _290__Convert_Binary_Number_in_a_Linked_List_to_Integer


/**

https://leetcode-cn.com/problems/convert-binary-number-in-a-linked-list-to-integer/

 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {

	num :=0
	if head!=nil{
		num<<=1
		num+=head.Val
		head=head.Next
	}

	return num
}
