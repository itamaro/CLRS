// Copyright 2018 Itamar Ostricher
// My MergeSort implementation for int-slices

package sort

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
