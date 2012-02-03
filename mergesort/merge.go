package mergesort

import (
	"sort"
)

// Mergesort the simple in place version...
func MergeSort(data sort.Interface) { mergeSort(data, 0, data.Len()) }

func mergeSort(data sort.Interface, a, b int) {
	if b <= a {
		return
	}
	mid := a + (b-a)/2
	mergeSort(data, a, mid)
	mergeSort(data, mid+1, b)

	merge(data, a, mid, mid+1, b)
}

func merge(data sort.Interface, a, b, c, d int) {
	// TODO in place merge is a mess. will do soon
}

// Insertion sort - swiped from the sort package since it was
// internal.
func insertionSort(data sort.Interface, a, b int) {
	for i := a + 1; i < b; i++ {
		for j := i; j > a && data.Less(j, j-1); j-- {
			data.Swap(j, j-1)
		}
	}
}
