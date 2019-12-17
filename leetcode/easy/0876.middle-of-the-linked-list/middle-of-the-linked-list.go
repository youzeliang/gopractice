package _876_middle_of_the_linked_list

/**

 https://leetcode-cn.com/problems/middle-of-the-linked-list/

* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
