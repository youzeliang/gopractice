package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func get(head *ListNode, k int) *ListNode {
	former, latter := head, head

	for i := 0; i < k; i++ {
		former = former.Next
	}

	for former != nil {
		former, latter = former.Next,latter.Next
	}

	return latter
}
