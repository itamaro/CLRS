// Copyright 2018 Itamar Ostricher
// linkedlist package providing a singly LinkedList data structure
// for educational purposes (no need to replace the builtin container/list package)

package linkedlist

// SLinkedList struct is a singly linked list data structure
// It has O(1) insert & O(N) remove runtime at any position,
// but no random access to list items
type SLinkedList struct {
  Head *SListNode
  Len uint32
}

// A ListNode, holding a value, and pointers to next list item
type SListNode struct {
  Value ItemType
  Next *SListNode
}

// Insert a new list item with value `val` at the head of the list - O(1)
func (l *SLinkedList) Insert(value ItemType) *SListNode {
  newNode := &SListNode{value, l.Head}
  l.Head = newNode
  l.Len++
  return newNode
}

// Insert a new list item with value `val` immediately after an existing `item` - O(1)
// Doesn't verify that the existing item is in the list `l`
func (l *SLinkedList) InsertAfter(node *SListNode, value ItemType) *SListNode {
  newNode := &SListNode{value, node.Next}
  node.Next = newNode
  l.Len++
  return newNode
}

// Remove the `item` from the list `l` and return the value - O(N)
// Doesn't verify that the item is in the list `l`
func (l *SLinkedList) Remove(node *SListNode) ItemType {
  if l.Head == node {
    l.Head = node.Next
  } else {
    // locate the prev element
    var prev *SListNode
    for prev = l.Head; prev.Next != node; prev = prev.Next { }
    prev.Next = node.Next
  }
  l.Len--
  return node.Value
}

// Return a slice with the list item values in order from front to back - O(N)
func (l SLinkedList) AsSlice() []ItemType {
  slice := make([]ItemType, l.Len)
  for e, i := l.Head, 0; e != nil; e, i = e.Next, i + 1 {
    slice[i] = e.Value
  }
  return slice
}
