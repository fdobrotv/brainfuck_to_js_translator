package stack

import (
	"errors"
)

type (
	Stack struct {
		top *Node
	}

	Node struct {
		value uint
		prev  *Node
	}
)

func New() *Stack {
	return &Stack{nil}
}

func (stack *Stack) Pop() (uint, error) {
	if stack.top == nil {
		return 0, errors.New("empty stack")
	}

	node := stack.top
	stack.top = node.prev

	return node.value, nil
}

func (stack *Stack) Push(value uint) {
	node := &Node{value, stack.top}
	stack.top = node
}
