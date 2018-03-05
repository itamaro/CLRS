// Copyright 2018 Itamar Ostricher
// My QuickSort implementation for int-slices

package sort

// Sort the given `slice` in-place (non-decreasing order) using the quick sort algorithm
// Runs in O(NlgN) up to O(N^2) time with up to O(N) memory (recursion)
func QuickSort(slice []int) {
	if len(slice) < 2 {
		return
	}
	p := partition(slice)
	QuickSort(slice[:p])
	QuickSort(slice[p+1:])
}

// Partition the slice
// Choose the right-most element as the pivot,
// and move all elements that are smaller to the left of it,
// such that eventually the pivot ends up in its final position
func partition(slice []int) int {
	left, right := 0, len(slice) - 1
	for pivot, i := slice[right], 0; i < right; i++ {
		if slice[i] < pivot {
			if i != left {
				slice[left], slice[i] = slice[i], slice[left]
			}
			left++
		}
	}
	slice[left], slice[right] = slice[right], slice[left]
	return left
}
