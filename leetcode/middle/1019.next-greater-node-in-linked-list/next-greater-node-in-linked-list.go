package _019_next_greater_node_in_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func nextLargerNodes(head *ListNode) []int {
	A := convert(head)
	res := make([]int, len(A))
	stack, peek := make([]int, len(A)), -1
	for i := 0; i < len(A); i++ {
		for peek >= 0 && A[stack[peek]] < A[i] {
			res[stack[peek]] = A[i]
			peek--
		}
		peek++
		stack[peek] = i
	}
	return res
}

func convert(head *ListNode) []int {
	res := make([]int, 0, 1024)
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}
