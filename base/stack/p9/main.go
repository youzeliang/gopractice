package main

import (
	"errors"
	"fmt"
)

type Element interface{}

type Stack struct {
	elements []Element
	top      int // 栈顶指针
	cap      int // 容量
}

// Push 作用：Push 方法就是往stack的存储区域压入新的元素
func (stack *Stack) Push(element Element) (err error) {
	// top == cap时，栈满
	if stack.top >= stack.cap {
		return errors.New("the stack is full")
	}
	stack.elements[stack.top] = element
	stack.top++
	return nil
}

// Pop  作用：Pop 方法是取出栈顶元素，并且在存储区域内删除
func (stack *Stack) Pop() (ele Element, err error) {
	// top == 0时，栈空
	if stack.top <= 0 {
		return nil, errors.New("the stack is empty")
	}
	ele = stack.elements[stack.top]
	// 在栈中清除当前元素
	stack.elements = append(stack.elements, stack.elements[:stack.top], stack.elements[stack.top+1:])
	stack.top--
	return ele, nil
}

// Peek  作用：Peek 方法就是返回栈顶的值，但是不删除存储区域的元素
func (stack *Stack) Peek() (ele Element, err error) {
	if stack.top <= 0 {
		return nil, errors.New("the stack is empty")
	}
	ele = stack.elements[stack.top]
	return ele, nil
}

// 初始化栈  一个创建栈的方法：NewStack
func NewStack(cap int) *Stack {
	elements := make([]Element, cap)
	return &Stack{
		elements: elements,
		top:      0,
		cap:      cap,
	}
}

func main() {
	stack := NewStack(5)
	for i := 0; i < 4; i++ {
		var ele Element = i
		stack.Push(ele)
	}

	element, _ := stack.Pop()
	fmt.Println(element)

	peekEle, _ := stack.Peek()
	fmt.Println(peekEle)

	fmt.Println(stack.Len())
	fmt.Println(stack.Cap())

	stack.Clear()
}
