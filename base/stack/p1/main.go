package main

import "fmt"

// go 语言写一个栈
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

func (s *Stack) Peek() int {
	if s.IsEmpty() {
		panic("Stack is empty")
	}
	return s.data[len(s.data)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack) Size() int {
	return len(s.data)
}

func main() {
	stack := Stack{}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	fmt.Println("Stack Size:", stack.Size())
	fmt.Println("Top Element:", stack.Peek())

	fmt.Println("Popping elements:")
	for !stack.IsEmpty() {
		fmt.Println(stack.Pop())
	}

	fmt.Println("Stack Size:", stack.Size())
}
