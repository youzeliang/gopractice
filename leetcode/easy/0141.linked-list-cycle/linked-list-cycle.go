package _141_linked_list_cycle

// https://leetcode-cn.com/problems/linked-list-cycle/
//
//思路
//
//快指针一次走两步，慢指针一次走一步，如果链表有环，那么两个指针始终会相遇。

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow, fast := head, head.Next

	for fast != nil && fast.Next != nil && slow != fast {
		slow, fast = slow.Next, fast.Next.Next
	}

	return slow == fast
}
