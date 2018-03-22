// Copyright 2018 Itamar Ostricher
// My BubbleSort implementation for int-slices

package sort

// Sort the given `slice` in-place (non-decreasing order) using the bubble sort algorithm
// Runs in O(N^2) time with no extra memory
func BubbleSort(slice []int) {
    for i := 0; i < len(slice); i++ {
        for j := 1; j < len(slice); j++ {
            if slice[j] < slice[j - 1] {
                slice[j - 1], slice[j] = slice[j], slice[j - 1]
            }
        }
    }
}
