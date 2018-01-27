// Copyright 2018 Itamar Ostricher
// queue package providing a Queue data structure
// for educational purposes (quite similar to github.com/golang-collections/collections/queue)

package queue

// Queue struct provides queue semantics data structure,
// with O(1) enqueue/dequeue/peek/isempty operations
type Queue struct {
	front, back *queueNode
}

// internal structure for queue nodes
type queueNode struct {
	Value ItemType
	next *queueNode
}

// Type of queue item values
type ItemType interface{}

// Enqueue a new item at the end of the queue
// Runs in O(1) with no extra memory
func (q *Queue) Enqueue(value ItemType) {
	e := &queueNode{value, nil}
	if q.back == nil {
		q.front, q.back = e, e
	} else {
		q.back.next = e
		q.back = e
	}
}

// Dequeue an item from beginning of the queue
// Runs in O(1) with no extra memory
func (q *Queue) Dequeue() (v ItemType) {
	if q.front == nil {
		panic("Dequeue from empty queue!")
	}
	v = q.front.Value
	if q.front == q.back {
		q.front, q.back = nil, nil
	} else {
		q.front = q.front.next
	}
	return
}

// Peek at the item at the beginning of the queue without dequeueing it
// Runs in O(1) with no extra memory
func (q Queue) Peek() ItemType {
	if q.front == nil {
		panic("Peek with empty queue!")
	}
	return q.front.Value
}

// Return true if the queue is empty, false otherwise
// Runs in O(1) with no extra memory
func (q Queue) IsEmpty() bool {
	return q.front == nil
}
