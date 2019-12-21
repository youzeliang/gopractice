package _876_middle_of_the_linked_list

/**

 https://leetcode-cn.com/problems/middle-of-the-linked-list/

* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
*/

/**
解题思路

定义一个快指针fast 一个慢指针slow ，快指针一次移动两个结点，慢指针一次移动一个结点

当fast到达链表的尾部结点时，慢指针也就移动到了链表的中间结点（同化成一个路程问题，同一段路程，A的速度是B的两倍，他们同时出发，当A走完全程时，B也就刚好走过一半）

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
