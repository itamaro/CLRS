// Copyright 2018 Itamar Ostricher

package stack

import "testing"

func TestStack(t *testing.T) {
	var s Stack
	if !s.IsEmpty() {
		t.Errorf("Expected new stack to be empty...")
	}
	for i := 0; i < 10; i++ {
		s.Push(i)
	}
	if s.IsEmpty() {
		t.Errorf("Expected stack after push to be non-empty...")
	}
	if v := s.Peek().(int); v != 9 {
		t.Errorf("Expected Peek to return 9, got %d", v)
	}
	for i := 9; i >= 0; i-- {
		n := s.Pop()
		if n.(int) != i {
			t.Errorf("Unexpected popped value %d - expected %d", n.(int), i)
		}
	}
	if !s.IsEmpty() {
		t.Errorf("Expected popped stack to be empty...")
	}
}
