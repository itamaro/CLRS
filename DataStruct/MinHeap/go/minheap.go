// Copyright 2018 Itamar Ostricher
// minheap package providing an int MinHeap data structure
// for educational purposes

package minheap

// MinHeap struct is a heap binary tree (in slice) such that the minimum value is always at the root
// provides Add / Peek / Pop semantics
type MinHeap struct {
  data []int
  length int
}

// Add new value to a heap - O(logN)
func (heap *MinHeap) Add(value int) {
  if heap.length == len(heap.data) {
    heap.resize(2 * len(heap.data))
  }
  heap.data[heap.length] = value
  heap.length++
  heap.bubbleUp()
}

// Return current heap length (number of elements) - O(1)
func (heap MinHeap) Len() int {
  return heap.length
}

// Return minimal element in the heap - O(1)
func (heap MinHeap) Peek() int {
  if len(heap.data) > 0 {
    return heap.data[0]
  }
  panic("Heap is empty!")
}

// Remove the minimal element in the heap and return its value - O(logN)
func (heap *MinHeap) Pop() (value int) {
  if len(heap.data) > 0 {
    value = heap.data[0]
    heap.length--
    heap.data[0] = heap.data[heap.length]
    heap.bubbleDown()
    if heap.length < len(heap.data) / 3 {
      heap.resize(len(heap.data) / 2)
    }
    return
  }
  panic("Heap is empty!")
}

// Grow (or shrink) underlying data slice
func (heap *MinHeap) resize(size int) {
  if size < 8 {
    size = 8
  }
  if size > len(heap.data) {
    newData := make([]int, size)
    copy(newData, heap.data)
    heap.data = newData
  } else if size < len(heap.data) {
    heap.data = heap.data[:size]
  }
}

// Bubble up the last element to its correct position
func (heap MinHeap) bubbleUp() {
  for cur, parent := heap.length - 1, (heap.length - 2) / 2;
      cur > 0 && heap.data[cur] < heap.data[parent];
      cur, parent = parent, (parent - 1) / 2 {
    heap.data[cur], heap.data[parent] = heap.data[parent], heap.data[cur]
  }
}

// Bubble down the first element to its correct position
func (heap MinHeap) bubbleDown() {
  i := 0
  for {
    cur := heap.data[i]
    c, minChild := heap.getMinChild(i)
    if c >= 0 && cur > minChild {
      heap.data[i], heap.data[c] = minChild, cur
      i = c
    } else {
      return
    }
  }
}

// Return the index and value of the minimal child of node at position i
// Returns index -1 if there are no children
func (heap MinHeap) getMinChild(i int) (j, val int) {
  j = -1
  if l := 2 * i + 1; l < heap.length {
    j, val = l, heap.data[l]
  }
  if r := 2 * i + 2; r < heap.length && heap.data[r] < val {
    j, val = r, heap.data[r]
  }
  return
}
