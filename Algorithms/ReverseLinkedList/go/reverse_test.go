// Copyright 2018 Itamar Ostricher

package main

import (
  linkedlist "github.com/itamaro/CLRS/DataStruct/LinkedList/go"
  "reflect"
  "testing"
)

func TestReverseList(t *testing.T) {
	var l linkedlist.LinkedList
	for i := 0; i < 1000; i++ {
		l.PushBack(i)
	}
  front, back := l.Front, l.Back
  slice := l.AsReversedSlice()
	Reverse(&l)
	if l.Len != 1000 {
		t.Errorf("Expected reversed list to have length 1000, got %d", l.Len)
	}
  if l.Front != back || l.Back != front {
    t.Errorf("Expected list Front & Back to switch roles")
  }
  if !reflect.DeepEqual(l.AsSlice(), slice) {
    t.Errorf("Expected list slice to be reversed")
  }
}
