package _328_odd_even_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	// oe: odd end 指向 odd group 的最后一个 node
	oe := head
	// eb: even begin 指向 even group 的第一个 node
	eb := head.Next
	// ee: even end 指向 even group 的最后一个 node
	ee := eb

	for ee != nil && ee.Next != nil {
		oe.Next = ee.Next
		oe = oe.Next
		ee.Next = oe.Next
		ee = ee.Next
		oe.Next = eb
	}

	return head
}
