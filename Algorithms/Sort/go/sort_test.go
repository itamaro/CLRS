// Copyright 2018 Itamar Ostricher

package sort

import (
    "math/rand"
    "reflect"
    "sort"
    "testing"
)

type testSpec struct {
  arr []int
  sorted []int
}

var TEST_CASES = [...]testSpec{
    testSpec{[]int{0, 5, 6, 1, 9, 2, 0, 3, 5, 4, 7, 6},
             []int{0, 0, 1, 2, 3, 4, 5, 5, 6, 6, 7, 9}},
    testSpec{[]int{1, 2, 3},
             []int{1, 2, 3}},
    testSpec{[]int{},
             []int{}},
    testSpec{[]int{-100},
             []int{-100}},
    testSpec{[]int{1, 1, 1},
             []int{1, 1, 1}},
    testSpec{[]int{5, 4, 3, 2, 1},
             []int{1, 2, 3, 4, 5}},
}

func TestInsertionSort(t *testing.T) {
  for _, testCase := range TEST_CASES {
    arr, sorted := testCase.arr, testCase.sorted
    InsertionSort(arr)
    if !reflect.DeepEqual(arr, sorted) {
      t.Errorf("Expected sorted array, got %v", arr)
    }
  }
}

func TestBubbleSort(t *testing.T) {
  for _, testCase := range TEST_CASES {
    arr, sorted := testCase.arr, testCase.sorted
    BubbleSort(arr)
    if !reflect.DeepEqual(arr, sorted) {
      t.Errorf("Expected sorted array, got %v", arr)
    }
  }
}

func TestSelectionSort(t *testing.T) {
  for _, testCase := range TEST_CASES {
    arr, sorted := testCase.arr, testCase.sorted
    SelectionSort(arr)
    if !reflect.DeepEqual(arr, sorted) {
      t.Errorf("Expected sorted array, got %v", arr)
    }
  }
}

func TestMergeSort(t *testing.T) {
  for _, testCase := range TEST_CASES {
    arr, sorted := testCase.arr, testCase.sorted
    MergeSort(arr)
    if !reflect.DeepEqual(arr, sorted) {
      t.Errorf("Expected sorted array, got %v", arr)
    }
  }
}

func TestGoMergeSort(t *testing.T) {
  for _, testCase := range TEST_CASES {
    arr, sorted := testCase.arr, testCase.sorted
    GoMergeSort(arr)
    if !reflect.DeepEqual(arr, sorted) {
      t.Errorf("Expected sorted array, got %v", arr)
    }
  }
}

func TestChanMergeSort(t *testing.T) {
  for _, testCase := range TEST_CASES {
    arr, sorted := testCase.arr, testCase.sorted
    ChanMergeSort(arr)
    if !reflect.DeepEqual(arr, sorted) {
      t.Errorf("Expected sorted array, got %v", arr)
    }
  }
}

func TestQuickSort(t *testing.T) {
  for _, testCase := range TEST_CASES {
    arr, sorted := testCase.arr, testCase.sorted
    QuickSort(arr)
    if !reflect.DeepEqual(arr, sorted) {
      t.Errorf("Expected sorted array, got %v", arr)
    }
  }
}

const BENCH_N int = 10000

func BenchmarkInsertionSort(b *testing.B) {
  for i := 0; i < b.N; i++ {
    // don't account for the creation of a random permutation
    b.StopTimer()
    slice := rand.Perm(BENCH_N)
    b.StartTimer()
    // measure just the sort
    InsertionSort(slice)
  }
}

func BenchmarkBubbleSort(b *testing.B) {
  for i := 0; i < b.N; i++ {
    // don't account for the creation of a random permutation
    b.StopTimer()
    slice := rand.Perm(BENCH_N)
    b.StartTimer()
    // measure just the sort
    BubbleSort(slice)
  }
}

func BenchmarkSelectionSort(b *testing.B) {
  for i := 0; i < b.N; i++ {
    // don't account for the creation of a random permutation
    b.StopTimer()
    slice := rand.Perm(BENCH_N)
    b.StartTimer()
    // measure just the sort
    SelectionSort(slice)
  }
}

func BenchmarkMergeSort(b *testing.B) {
  for i := 0; i < b.N; i++ {
    // don't account for the creation of a random permutation
    b.StopTimer()
    slice := rand.Perm(BENCH_N)
    b.StartTimer()
    // measure just the sort
    MergeSort(slice)
  }
}

func BenchmarkGoMergeSort(b *testing.B) {
  for i := 0; i < b.N; i++ {
    // don't account for the creation of a random permutation
    b.StopTimer()
    slice := rand.Perm(BENCH_N)
    b.StartTimer()
    // measure just the sort
    GoMergeSort(slice)
  }
}

func BenchmarkChanMergeSort(b *testing.B) {
  for i := 0; i < b.N; i++ {
    // don't account for the creation of a random permutation
    b.StopTimer()
    slice := rand.Perm(BENCH_N)
    b.StartTimer()
    // measure just the sort
    ChanMergeSort(slice)
  }
}

func BenchmarkQuickSort(b *testing.B) {
  for i := 0; i < b.N; i++ {
    // don't account for the creation of a random permutation
    b.StopTimer()
    slice := rand.Perm(BENCH_N)
    b.StartTimer()
    // measure just the sort
    QuickSort(slice)
  }
}

func BenchmarkLibSort(b *testing.B) {
  for i := 0; i < b.N; i++ {
    // don't account for the creation of a random permutation
    b.StopTimer()
    slice := rand.Perm(BENCH_N)
    b.StartTimer()
    // measure just the sort
    sort.Slice(slice, func (i, j int) bool {
        return slice[i] < slice[j]
    })
  }
}
