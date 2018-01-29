// Copyright 2018 Itamar Ostricher

package sort

import (
    "reflect"
    "testing"
)

type testSpec struct {
  arr []int
  sorted []int
}

func TestInsertionSort(t *testing.T) {
  cases := []testSpec{
      testSpec{[]int{0, 5, 6, 1, 9, 2, 0, 3, 5, 4, 7, 6},
               []int{0, 0, 1, 2, 3, 4, 5, 5, 6, 6, 7, 9}},
      testSpec{[]int{1, 2, 3},
               []int{1, 2, 3}},
      testSpec{[]int{},
               []int{}},
      testSpec{[]int{1, 1, 1},
               []int{1, 1, 1}},
      testSpec{[]int{5, 4, 3, 2, 1},
               []int{1, 2, 3, 4, 5}},
  }
  for _, testCase := range cases {
    arr, sorted := testCase.arr, testCase.sorted
    InsertionSort(arr)
    if !reflect.DeepEqual(arr, sorted) {
      t.Errorf("Expected sorted array, got %v", arr)
    }
  }
}

func TestMergeSort(t *testing.T) {
  cases := []testSpec{
      testSpec{[]int{0, 5, 6, 1, 9, 2, 0, 3, 5, 4, 7, 6},
               []int{0, 0, 1, 2, 3, 4, 5, 5, 6, 6, 7, 9}},
      testSpec{[]int{1, 2, 3},
               []int{1, 2, 3}},
      testSpec{[]int{},
               []int{}},
      testSpec{[]int{1, 1, 1},
               []int{1, 1, 1}},
      testSpec{[]int{5, 4, 3, 2, 1},
               []int{1, 2, 3, 4, 5}},
  }
  for _, testCase := range cases {
    arr, sorted := testCase.arr, testCase.sorted
    MergeSort(arr)
    if !reflect.DeepEqual(arr, sorted) {
      t.Errorf("Expected sorted array, got %v", arr)
    }
  }
}
