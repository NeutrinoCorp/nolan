package list

import "github.com/neutrinocorp/nolan/collection"

// IndexOf returns the index of the given value. If not found, this routine will return -1.
func IndexOf[T comparable](list List[T], v T) int {
	iter := list.NewIterator()
	index := 0
	for iter.HasNext() {
		if val := iter.Next(); val == v {
			return index
		}
		index++
	}
	return -1
}

// IndexOfOrdered returns the index of the given value from an ordered List.
// If not found, this routine will return -1.
//
// It works as expected with slice/array implementations of List
// (e.g., SliceList) as this routine implements a binary search algorithm internally.
func IndexOfOrdered[T collection.Ordered](list List[T], v T) int {
	lowIdx := 0
	highIdx := list.Len() - 1

	for lowIdx <= highIdx {
		midIdx := lowIdx + (highIdx-lowIdx)/2
		elem := list.GetAt(midIdx)
		if elem == v {
			return midIdx
		} else if elem < v {
			lowIdx = midIdx + 1
			continue
		}

		highIdx = midIdx - 1
	}

	return -1
}
