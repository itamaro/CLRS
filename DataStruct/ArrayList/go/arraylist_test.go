// Copyright 2018 Itamar Ostricher

package arraylist

import "testing"

func TestArrayList(t *testing.T) {
  var myArr ArrayList
  if myArr.Len() != 0 {
    t.Errorf("Expected Len 0, got %d", myArr.Len())
  }
  myArr.Append(3)
  if myArr.Len() != 1 {
    t.Errorf("Expected Len 1, got %d", myArr.Len())
  }
  if myArr.Get(0) != 3 {
    t.Errorf("Expected myArr[0] to be 3, got %d", myArr.Get(0))
  }
  for i := 6; i < 100; i += 3 {
    myArr.Append(ItemType(i))
  }
  if myArr.Len() != 33 {
    t.Errorf("Expected Len 33, got %d", myArr.Len())
  }
  if myArr.Get(9) != 30 {
    t.Errorf("Expected myArr[9] to be 30, got %d", myArr.Get(9))
  }
  myArr.Set(9, ItemType(40))
  if myArr.Get(9) != 40 {
    t.Errorf("Expected myArr[9] to be 40, got %d", myArr.Get(9))
  }
  myArr.Set(9, ItemType(30))
  if myArr.capacity != 64 {
    t.Errorf("Expected myArr capacity to be 64, got %d", myArr.capacity)
  }
  for i := 99; i >= 3; i -= 3 {
    popped := myArr.Pop()
    if popped != ItemType(i) {
      t.Errorf("Expected myArr.Pop() to return %d, got %d", i, popped)
    }
  }
  if myArr.capacity != MIN_CAPACITY {
    t.Errorf("Expected myArr capacity to be %d, got %d", MIN_CAPACITY, myArr.capacity)
  }
}
