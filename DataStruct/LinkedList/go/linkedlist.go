// Copyright 2018 Itamar Ostricher
// linkedlist package providing a LinkedList data structure
// for educational purposes (no need to replace the builtin container/list package)

package linkedlist

// Item type for ListItem values
type ItemType interface{}

// LinkedList struct is a doubly linked list data structure
// It has O(1) insert / remove runtime at any position,
// but no random access to list items
type LinkedList struct {
  Len uint32  // get list length in O(1)
  Front *ListItem
  Back *ListItem
}

// A ListItem, holding a value, and pointers to prev/next list items
type ListItem struct {
  Value ItemType
  Next *ListItem
  Prev *ListItem
}

// Insert a new item with value `val` at the back of the list - O(1)
func (l *LinkedList) PushBack(val ItemType) *ListItem {
  newItem := &ListItem{val, nil, l.Back}
  l.Back = newItem
  l.Len++
  if l.Len == 1 {
    l.Front = newItem
  }
  if newItem.Prev != nil {
    newItem.Prev.Next = newItem
  }
  return newItem
}

// Insert a new item with value `val` at the front of the list - O(1)
func (l *LinkedList) PushFront(val ItemType) *ListItem {
  newItem := &ListItem{val, l.Front, nil}
  l.Front = newItem
  l.Len++
  if l.Len == 1 {
    l.Back = newItem
  }
  if newItem.Next != nil {
    newItem.Next.Prev = newItem
  }
  return newItem
}

// Insert a new list item with value `val` immediately after an existing `item` - O(1)
// Doesn't verify that the existing item is in the list `l`
func (l *LinkedList) InsertAfter(item *ListItem, val ItemType) *ListItem {
  if item == nil {
    panic("item can't be nil")
  }
  newItem := &ListItem{val, item.Next, item}
  item.Next = newItem
  if newItem.Next == nil {
    // so it's a PushBack
    l.Back = newItem
    // can't be also Front, because it's at least the second item
  } else {
    newItem.Next.Prev = newItem
  }
  l.Len++
  return newItem
}

// Insert a new list item with value `val` immediately before an existing `item` - O(1)
// Doesn't verify that the existing item is in the list `l`
func (l *LinkedList) InsertBefore(item *ListItem, val ItemType) *ListItem {
  if item == nil {
    panic("item can't be nil")
  }
  newItem := &ListItem{val, item, item.Prev}
  item.Prev = newItem
  if newItem.Prev == nil {
    // so it's a PushFront
    l.Front = newItem
    // can't be also Back, because it's at least the second item
  } else {
    newItem.Prev.Next = newItem
  }
  l.Len++
  return newItem
}

// Remove the `item` from the list `l` - O(1)
// Doesn't verify that the item is in the list `l`
func (l *LinkedList) Remove(item *ListItem) {
  if item == nil {
    return
  }
  if item.Prev == nil {
    // item is at front
    l.Front = item.Next
  } else {
    item.Prev.Next = item.Next
  }
  if item.Next == nil {
    // item is at back
    l.Back = item.Prev
  } else {
    item.Next.Prev = item.Prev
  }
  l.Len--
}

// Return a slice with the list item values in order from front to back - O(N)
func (l LinkedList) AsSlice() []ItemType {
  slice := make([]ItemType, l.Len)
  for item, i := l.Front, 0; item != nil; item, i = item.Next, i + 1 {
    slice[i] = item.Value
  }
  return slice
}

// Return a slice with the list item values in order from back to front - O(N)
func (l LinkedList) AsReversedSlice() []ItemType {
  slice := make([]ItemType, l.Len)
  for item, i := l.Back, 0; item != nil; item, i = item.Prev, i + 1 {
    slice[i] = item.Value
  }
  return slice
}
