// Copyright 2018 Itamar Ostricher

package queue

import "testing"

func TestQueue(t *testing.T) {
	var q Queue
	if !q.IsEmpty() {
		t.Errorf("Expected new queue to be empty...")
	}
	for i := 0; i < 10; i++ {
		q.Enqueue(i)
	}
	if q.IsEmpty() {
		t.Errorf("Expected used queue to be non-empty...")
	}
	if v := q.Peek().(int); v != 0 {
		t.Errorf("Expected Peek to return 0, got %v", v)
	}
	for i := 0; i < 10; i++ {
		v := q.Dequeue().(int)
		if v != i {
			t.Errorf("Expected Dequeue to return %d, got %d", i, v)
		}
	}
}
