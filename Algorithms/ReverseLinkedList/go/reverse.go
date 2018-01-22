// Copyright 2018 Itamar Ostricher
// Simple & efficient (doubly) linked list reversal algorithm
// for educational purposes

package main

import (
	linkedlist "github.com/itamaro/CLRS/DataStruct/LinkedList/go"
	"fmt"
)

// Reverse a doubly linked list, in place, in O(N) runtime with O(1) memory
func Reverse(l *linkedlist.LinkedList) {
  for item := l.Front; item != nil; item = item.Prev {
    item.Next, item.Prev = item.Prev, item.Next
  }
  l.Front, l.Back = l.Back, l.Front
}

func main() {
  var l linkedlist.LinkedList
  for i := 0; i < 10; i++ {
    l.PushBack(i)
  }
  fmt.Printf("List:     %v\n", l.AsSlice())
  Reverse(&l)
  fmt.Printf("Reversed: %v\n", l.AsSlice())
}
