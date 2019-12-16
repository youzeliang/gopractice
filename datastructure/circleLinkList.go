package main

// go语言判断单链表中是否存在环
// 判断单链表是否存在环是一个经典的快慢指针问题，一个每次走一步的指针，和一个每次走两步的指针，如果链表里有环的话，两个指针最终肯定会相遇，具体实现如下：

type Node struct {
	value    int
	nextNode *Node
}

func hasCycle(head *Node) bool {

	slow := head
	fast := head

	for fast != nil && fast.nextNode != nil {
		slow = slow.nextNode
		fast = fast.nextNode.nextNode
		if fast == slow {
			return true
		}
	}

	return false
}
