// Package stack - Basic implementation of a stack type
// that can hold values of any type
package stack

import "errors"

// Stack : A stack containing values of any type
type Stack []interface{}

// Len get length of stack
func (stack Stack) Len() int {
	return len(stack)
}

// Cap get capacity of stack
func (stack Stack) Cap() int {
	return cap(stack)
}

// IsEmpty is the stack empty?
func (stack Stack) IsEmpty() bool {
	return len(stack) == 0
}

// Push push a value into the stack
func (stack *Stack) Push(x interface{}) {
	*stack = append(*stack, x)
}

// Top get the top value from the stack WITHOUT removing it
func (stack Stack) Top() (interface{}, error) {
	if len(stack) == 0 {
		return nil, errors.New("Top() error - stack is empty")
	}
	return stack[len(stack)-1], nil
}

// Pop remove and get the top value from the stack
func (stack *Stack) Pop() (interface{}, error) {
	theStack := *stack
	if len(theStack) == 0 {
		return nil, errors.New("Pop() error - stack is empty")
	}
	poppedValue := theStack[len(theStack)-1]
	*stack = theStack[:len(theStack)-1]
	return poppedValue, nil
}
