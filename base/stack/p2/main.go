package main

import "fmt"

// 两个栈实现一个队列
type Stack struct {
	data []int
}

func (s *Stack) Push(val int) {
	s.data = append(s.data, val)
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		panic("Stack is empty")
	}
	length := len(s.data)
	val := s.data[length-1]
	s.data = s.data[:length-1]
	return val
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

type Queue struct {
	stack1 *Stack // 用于入队
	stack2 *Stack // 用于出队
}

func NewQueue() *Queue {
	return &Queue{
		stack1: &Stack{},
		stack2: &Stack{},
	}
}

func (q *Queue) Enqueue(val int) {
	q.stack1.Push(val)
}

func (q *Queue) Dequeue() int {
	if q.stack2.IsEmpty() {
		for !q.stack1.IsEmpty() {
			val := q.stack1.Pop()
			q.stack2.Push(val)
		}
	}
	return q.stack2.Pop()
}

func main() {
	queue := NewQueue()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	fmt.Println(queue.Dequeue()) // Output: 1
	fmt.Println(queue.Dequeue()) // Output: 2

	queue.Enqueue(4)
	fmt.Println(queue.Dequeue()) // Output: 3
	fmt.Println(queue.Dequeue()) // Output: 4
}
