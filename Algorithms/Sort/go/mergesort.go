// Copyright 2018 Itamar Ostricher
// My MergeSort implementation for int-slices

package sort

import "sync"

// Sort the given `slice` in-place (non-decreasing order) using the merge sort algorithm
// Runs in O(NlgN) time with O(N) memory
func MergeSort(slice []int) {
  if len(slice) < 2 {
    return
  }
  q := (len(slice) + 1) / 2
  left, right := slice[:q], slice[q:]
  MergeSort(left)
  MergeSort(right)
  merge(slice, q)
}

// Sort the given `slice` in-place (non-decreasing order) using the merge sort algorithm
// and taking advantage of goroutines to do recursive calls on partial slices in parallel
// Runs in O(NlgN) time with O(N) memory
func GoMergeSort(slice []int) {
  var wg sync.WaitGroup
  wg.Add(1)
  goMergeSortHelper(slice, &wg)
}

func goMergeSortHelper(slice []int, wg *sync.WaitGroup) {
  defer wg.Done()
  if len(slice) < 2 {
    return
  }
  q := (len(slice) + 1) / 2
  var subWg sync.WaitGroup
  subWg.Add(2)
  go goMergeSortHelper(slice[:q], &subWg)
  go goMergeSortHelper(slice[q:], &subWg)
  subWg.Wait()
  merge(slice, q)
}

func merge(slice []int, q int) {
  tmp := make([]int, len(slice))
  copy(tmp, slice)
  for i, j, n := 0, q, 0; n < len(slice); n++ {
    if tmp[i] < tmp[j] {
      slice[n] = tmp[i]
      i++
      if i == q {
        copy(slice[n+1:], tmp[j:])
        break
      }
    } else {
      slice[n] = tmp[j]
      j++
      if j == len(tmp) {
        copy(slice[n+1:], tmp[i:q])
        break
      }
    }
  }
}
