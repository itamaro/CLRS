// Copyright 2018 Itamar Ostricher

package main

import (
  linkedlist "github.com/itamaro/CLRS/DataStruct/LinkedList/go"
  "reflect"
  "testing"
)

func TestReverseList(t *testing.T) {
	var l linkedlist.LinkedList

  // test nil case
  Reverse(&l)
  if l.Front != nil || l.Back != nil || l.Len != 0 {
    t.Errorf("Reversing empty list did something weird!")
  }

  // test trivial case
  e := l.PushBack(0)
  Reverse(&l)
  if l.Front != e || l.Back != e || l.Len != 1 {
    t.Errorf("Reversing list with one node did something weird!")
  }

  // test useful case
	for i := 1; i < 1000; i++ {
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

func TestReverseSList(t *testing.T) {
  var l linkedlist.SLinkedList

  // test nil case
  SReverse(&l)
  if l.Head != nil || l.Len != 0 {
    t.Errorf("Reversing empty list did something weird!")
  }

  // test trivial case
  e := l.Insert(0)
  SReverse(&l)
  if l.Head != e || l.Len != 1 {
    t.Errorf("Reversing list with one node did something weird!")
  }

  // test useful case
  for i := 1; i < 1000; i++ {
    e = l.InsertAfter(e, i)
  }
  SReverse(&l)
  if l.Len != 1000 {
    t.Errorf("Expected reversed list to have length 1000, got %d", l.Len)
  }
  for i, e := 999, l.Head; i >= 0; i, e = i - 1, e.Next {
    if i != e.Value.(int) {
      t.Fatalf("List not reversed")
    }
  }
}
