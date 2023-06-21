package _095

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	dummyHead := &ListNode{Next: head}
	pre, slow, fast := dummyHead, head, head
	for fast != nil && fast.Next != nil {
		pre = slow // pre 记录了 slow 的上一个结点
		slow = slow.Next
		fast = fast.Next.Next
	}
	pre.Next = slow.Next // 循环结束后 slow 为待删除节点
	return dummyHead.Next
}
