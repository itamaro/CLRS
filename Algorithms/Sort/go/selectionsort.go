// Copyright 2018 Itamar Ostricher
// My SelectionSort implementation for int-slices

package sort

// Sort the given `slice` in-place (non-decreasing order) using the selection sort algorithm
// Runs in O(N^2) time with no extra memory
func SelectionSort(slice []int) {
    for i := 0; i < len(slice) - 1; i++ {
        minIndex := i
        for j := i + 1; j < len(slice); j++ {
            if slice[j] < slice[minIndex] {
                minIndex = j
            }
        }
        if i != minIndex {
            slice[i], slice[minIndex] = slice[minIndex], slice[i]
        }
    }
}
