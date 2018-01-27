// Copyright 2018 Itamar Ostricher
// stack package providing a Stack data structure
// for educational purposes (quite similar to github.com/golang-collections/collections/stack)

package stack

// Stack struct provides stack semantics data structure,
// with O(1) push/pop/peek/isempty operations
type Stack struct {
	top *stackNode
}

// internal structure for stack nodes
type stackNode struct {
	value ItemType
	next *stackNode
}

// Type of stack item values
type ItemType interface{}

// Push a new item to the top of the stack
// Runs in O(1) with no extra memory
func (s *Stack) Push(value ItemType) {
	e := &stackNode{value, s.top}
	s.top = e
}

// Pop the top item from the stack
// Runs in O(1) with no extra memory
func (s *Stack) Pop() (v ItemType) {
	if s.top == nil {
		panic("Pop from empty stack!")
	}
	v = s.top.value
	s.top = s.top.next
	return
}

// Peek at the top item of the stack without popping it
// Runs in O(1) with no extra memory
func (s Stack) Peek() ItemType {
	if s.top == nil {
		panic("Peek with empty stack!")
	}
	return s.top.value
}

// Return true if the stack is empty, false otherwise
// Runs in O(1) with no extra memory
func (s Stack) IsEmpty() bool {
	return s.top == nil
}
