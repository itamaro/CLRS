// Copyright 2018 Itamar Ostricher

package minheap

import (
  "math/rand"
  "testing"
)

func TestMinHeapSortedAdd(t *testing.T) {
  var heap MinHeap
  for i := 0; i < 10; i++ {
    heap.Add(i)
  }
  if l := heap.Len(); l != 10 {
    t.Errorf("Expected heap length 10, got %d\n", l)
  }
  if top := heap.Peek(); top != 0 {
    t.Errorf("Expected heap top to be 0, got %d", top)
  }
  for i := 0; i < 10; i++ {
    if top := heap.Pop(); top != i {
      t.Fatalf("Expected heap top to be %d, got %d", i, top)
    }
    if l := heap.Len(); l != 9 - i {
      t.Fatalf("Expected heap length %d, got %d", 9 - i, l)
    }
  }
}

func TestMinHeapRand(t *testing.T) {
  var heap MinHeap
  for i := 0; i < 1000; i++ {
    heap.Add(rand.Int())
  }
  cur := heap.Pop()
  for heap.Len() > 0 {
    if next := heap.Pop(); next < cur {
      t.Fatalf("Next heap value %d smaller than current %d", next, cur)
    } else {
      cur = next
    }
  }
}
