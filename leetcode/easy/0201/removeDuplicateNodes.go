package _201

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeDuplicateNodes(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	m := map[int]int{}
	for cur := dummy; cur != nil && cur.Next != nil; {
		nextV := cur.Next.Val
		if _, ok := m[nextV]; ok {
			cur.Next = cur.Next.Next
		} else {
			m[nextV] = 1
			cur = cur.Next
		}
	}
	return dummy.Next

}
