// Copyright 2018 Itamar Ostricher

package linkedlist

import (
  "reflect"
  "testing"
)

func TestLinkedList(t* testing.T) {
  var l LinkedList
  if l.Len != 0 {
    t.Errorf("Expected new list length to be 0, got %d", l.Len)
  }
  if l.Front != nil || l.Back != nil {
    t.Errorf("Expected new list to have nil Front & Back")
  }
  n1 := l.PushBack("first")
  if l.Len != 1 {
    t.Errorf("Expected list length to be 1, got %d", l.Len)
  }
  if l.Front != n1 || l.Back != n1 {
    t.Errorf("Expected one-item list to have Front & Back equal to the item")
  }
  if n1.Value != "first" {
    t.Errorf("Expected added item to have value 'first', got %v", n1.Value)
  }
  n2 := l.PushBack("second")
  if l.Front != n1 {
    t.Errorf("Expected list front to be %v, got %v", n1, l.Front)
  }
  if l.Back != n2 {
    t.Errorf("Expected list back to be %v, got %v", n2, l.Back)
  }
  if n2.Value != "second" {
    t.Errorf("Expected added item to have value 'second', got %v", n2.Value)
  }
  n3 := l.PushFront("third")
  if l.Front != n3 {
    t.Errorf("Expected list front to be %v, got %v", n3, l.Front)
  }
  if l.Back != n2 {
    t.Errorf("Expected list back to be %v, got %v", n2, l.Back)
  }
  if n3.Value != "third" {
    t.Errorf("Expected added item to have value 'third', got %v", n3.Value)
  }
  if l.Len != 3 {
    t.Errorf("Expected list length to be 3, got %d", l.Len)
  }
  if n3.Prev != nil || n3.Next != n1 {
    t.Errorf("Expected front item to have nil Prev and %v as next, got %v", n1, n3)
  }
  if n1.Prev != n3 || n1.Next != n2 {
    t.Errorf("Expected middle item to have %v as Prev and %v as next, got %v", n3, n2, n1)
  }
  if n2.Prev != n1 || n2.Next != nil {
    t.Errorf("Expected back item to have %v as Prev and nil next, got %v", n1, n2)
  }
  if !reflect.DeepEqual(l.AsSlice(), []ItemType{"third", "first", "second"}) {
    t.Errorf("Expected list as slice to be third-first-second, got %v", l.AsSlice())
  }
  if !reflect.DeepEqual(l.AsReversedSlice(), []ItemType{"second", "first", "third"}) {
    t.Errorf("Expected list as reversed slice to be second-first-third, got %v", l.AsReversedSlice())
  }
  l.InsertAfter(n1, "fourth")
  if l.Len != 4 {
    t.Errorf("Expected list length to be 4, got %d", l.Len)
  }
  if !reflect.DeepEqual(l.AsSlice(), []ItemType{"third", "first", "fourth", "second"}) {
    t.Errorf("Expected list as slice to be third-first-fourth-second, got %v", l.AsSlice())
  }
  n5 := l.InsertAfter(n2, "fifth")
  if l.Len != 5 {
    t.Errorf("Expected list length to be 5, got %d", l.Len)
  }
  if !reflect.DeepEqual(l.AsSlice(), []ItemType{"third", "first", "fourth", "second", "fifth"}) {
    t.Errorf("Expected list as slice to be third-first-fourth-second-fifth, got %v", l.AsSlice())
  }
  l.InsertBefore(n2, "sixth")
  if l.Len != 6 {
    t.Errorf("Expected list length to be 6, got %d", l.Len)
  }
  if !reflect.DeepEqual(l.AsSlice(), []ItemType{"third", "first", "fourth", "sixth", "second", "fifth"}) {
    t.Errorf("Expected list as slice to be third-first-fourth-sixth-second-fifth, got %v", l.AsSlice())
  }
  n7 := l.InsertBefore(n3, "seventh")
  if l.Len != 7 {
    t.Errorf("Expected list length to be 7, got %d", l.Len)
  }
  if !reflect.DeepEqual(l.AsSlice(), []ItemType{"seventh", "third", "first", "fourth", "sixth", "second", "fifth"}) {
    t.Errorf("Expected list as slice to be seventh-third-first-fourth-sixth-second-fifth, got %v", l.AsSlice())
  }
  // test remove from middle
  l.Remove(n1)
  l.Remove(n2)
  l.Remove(n3)
  // as well as from front & back
  l.Remove(n7)
  l.Remove(n5)
  if l.Len != 2 {
    t.Errorf("Expected list length to be 2, got %d", l.Len)
  }
  if !reflect.DeepEqual(l.AsSlice(), []ItemType{"fourth", "sixth"}) {
    t.Errorf("Expected list as slice to be fourth-sixth, got %v", l.AsSlice())
  }
}

func TestSLinkedList(t *testing.T) {
  var l SLinkedList
  if l.Len != 0 {
    t.Errorf("Expected new list length to be 0, got %d", l.Len)
  }
  if l.Head != nil {
    t.Errorf("Expected new list to have nil Head")
  }
  n1 := l.Insert("first")
  n2 := l.Insert("second")
  if l.Len != 2 {
    t.Errorf("Expected list length to be 2, got %d", l.Len)
  }
  if n1.Value != "first" {
    t.Errorf("Expected added item to have value 'first', got %v", n1.Value)
  }
  if n2.Value != "second" {
    t.Errorf("Expected added item to have value 'second', got %v", n2.Value)
  }
  if l.Head != n2 || n2.Next != n1 || n1.Next != nil {
    t.Errorf("List structure could not be verified")
  }
  n3 := l.InsertAfter(n1, "third")
  l.InsertAfter(n2, "fourth")
  if l.Len != 4 {
    t.Errorf("Expected list length to be 4, got %d", l.Len)
  }
  if !reflect.DeepEqual(l.AsSlice(), []ItemType{"second", "fourth", "first", "third"}) {
    t.Errorf("Expected list as slice to be second-fourth-first-third, got %v", l.AsSlice())
  }
  // remove from middle
  if l.Remove(n1) != "first" {
    t.Errorf("Expected Remove to return the value of the removed node")
  }
  if l.Len != 3 {
    t.Errorf("Expected list length to be 3, got %d", l.Len)
  }
  if !reflect.DeepEqual(l.AsSlice(), []ItemType{"second", "fourth", "third"}) {
    t.Errorf("Expected list as slice to be second-fourth-third, got %v", l.AsSlice())
  }
  // remove from end
  l.Remove(n3)
  if l.Len != 2 {
    t.Errorf("Expected list length to be 2, got %d", l.Len)
  }
  if !reflect.DeepEqual(l.AsSlice(), []ItemType{"second", "fourth"}) {
    t.Errorf("Expected list as slice to be second-fourth, got %v", l.AsSlice())
  }
  // remove from head
  l.Remove(n2)
  if l.Len != 1 {
    t.Errorf("Expected list length to be 1, got %d", l.Len)
  }
  if !reflect.DeepEqual(l.AsSlice(), []ItemType{"fourth"}) {
    t.Errorf("Expected list as slice to be fourth, got %v", l.AsSlice())
  }
}
