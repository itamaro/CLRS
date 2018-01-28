// Copyright 2018 Itamar Ostricher
// My InsertSort implementation for int-slices

package sort

// Sort the given `slice` in-place (non-decreasing order) using the insertion sort algorithm
// Runs in O(N^2) time with no extra memory
func InsertionSort(slice []int) {
  for i := 1; i < len(slice); i++ {
    for j := i; j > 0 && slice[j] < slice[j - 1]; j-- {
      slice[j - 1], slice[j] = slice[j], slice[j - 1]
    }
  }
}
