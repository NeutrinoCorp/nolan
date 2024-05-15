package list

import (
	"github.com/neutrinocorp/nolan/collection"
)

// SliceList is the Go's slice implementation of List. The user of this interface has precise control
// over where in the list each element is inserted. The user can access elements by their integer
// index (position in the list), and search for elements in the list.
type SliceList[T any] struct {
	Source []T
}

var (
	_ List[int] = &SliceList[int]{}
)

// NewSliceList allocates a new SliceList instance.
func NewSliceList[T any](src []T) *SliceList[T] {
	if src == nil {
		src = make([]T, 0)
	}
	return &SliceList[T]{
		Source: src,
	}
}

// NewSliceListFromCollection allocates a new SliceList instance copying 'src' items into it.
func NewSliceListFromCollection[T any](src collection.Collection[T]) *SliceList[T] {
	ls := &SliceList[T]{}
	iter := src.NewIterator()
	for iter.HasNext() {
		ls.Add(iter.Next())
	}
	return ls
}

func (s *SliceList[T]) growIfRequired(n int) {
	if cap(s.Source)-len(s.Source) >= n {
		return
	}
	s.Source = append(make([]T, 0, len(s.Source)+n), s.Source...)
}

// NewIterator returns an iterator over elements of type T.
func (s *SliceList[T]) NewIterator() collection.Iterator[T] {
	return NewIterator[T](s)
}

// Add adds an element into this collection.
func (s *SliceList[T]) Add(v T) bool {
	s.Source = append(s.Source, v)
	return true
}

// AddAll adds all the elements into this collection.
func (s *SliceList[T]) AddAll(src collection.Collection[T]) bool {
	s.growIfRequired(src.Len())
	wasMod := false
	src.ForEach(func(a T) bool {
		wasAdded := s.Add(a)
		if wasAdded && !wasMod {
			wasMod = true
		}
		return false
	})
	return wasMod
}

// AddSlice adds all the elements in the specified slice (variadic) to this collection.
func (s *SliceList[T]) AddSlice(items ...T) bool {
	if len(items) == 0 {
		return false
	}
	s.growIfRequired(len(items))
	s.Source = append(s.Source, items...)
	return true
}

// Clear Removes all the elements from this collection. Does not de-allocates Source.
func (s *SliceList[T]) Clear() {
	s.Source = s.Source[:0]
}

// Len returns the number of elements in this collection.
func (s *SliceList[T]) Len() int {
	return len(s.Source)
}

// IsEmpty returns true if this collection contains no elements.
func (s *SliceList[T]) IsEmpty() bool {
	return len(s.Source) == 0
}

// ToSlice returns all the elements from this collection as a slice of T.
func (s *SliceList[T]) ToSlice() []T {
	if len(s.Source) == 0 {
		// We return a nil slice instead the underlying slice here as factory func (NewSliceList)
		// allocates a slice if nil is passed as src.
		// This leads to a never-nil return in this routine and thus, break list API homogeneity.
		return nil
	}
	return s.Source
}

// ForEach traverses through all the elements from this collection.
// Use predicate's return value to indicate a break of the iteration, TRUE meaning a break.
func (s *SliceList[T]) ForEach(predicateFunc collection.IterablePredicateFunc[T]) {
	for _, item := range s.Source {
		willBreak := predicateFunc(item)
		if willBreak {
			break
		}
	}
}

// AddAt inserts the specified element at the specified position in this list.
func (s *SliceList[T]) AddAt(index int, v T) {
	if !isValidIndex(index, len(s.Source)) {
		return
	}

	insertionIndex := index + 1
	s.growIfRequired(1)
	s.Source = append(s.Source[:insertionIndex], append([]T{v}, s.Source[insertionIndex:]...)...)
}

// AddAllAt inserts all the elements in the specified collection into this list at the specified position.
func (s *SliceList[T]) AddAllAt(index int, src collection.Collection[T]) bool {
	if !isValidIndex(index, len(s.Source)) {
		return false
	}
	insertionIndex := index + 1
	s.growIfRequired(src.Len())
	newSlice := src.ToSlice()
	s.Source = append(s.Source[:insertionIndex], append(newSlice, s.Source[insertionIndex:]...)...)
	return true
}

// SetAt replaces the element at the specified position in this list with the specified element.
func (s *SliceList[T]) SetAt(index int, v T) T {
	if !isValidIndex(index, len(s.Source)) {
		var zeroVal T
		return zeroVal
	}

	tmpVal := s.Source[index]
	s.Source[index] = v
	return tmpVal
}

// GetAt returns the element at the specified position in this list.
func (s *SliceList[T]) GetAt(index int) T {
	if !isValidIndex(index, len(s.Source)) {
		var zeroVal T
		return zeroVal
	}
	return s.Source[index]
}

// RemoveAt removes the element at the specified position in this list.
func (s *SliceList[T]) RemoveAt(index int) T {
	if !isValidIndex(index, len(s.Source)) {
		var zeroVal T
		return zeroVal
	}
	item := s.GetAt(index)
	copy(s.Source[index:], s.Source[index+1:])
	var zeroVal T
	s.Source[len(s.Source)-1] = zeroVal
	s.Source = s.Source[:len(s.Source)-1]
	return item
}

// ForEachWithIndex traverses through all the elements from this collection.
// Use predicate's return value to indicate a break of the iteration.
// 'A' is the index while 'B' is the item.
func (s *SliceList[T]) ForEachWithIndex(predicateFunc collection.IterablePredicateBiFunc[int, T]) {
	for i, item := range s.Source {
		willBreak := predicateFunc(i, item)
		if willBreak {
			break
		}
	}
}

// ToSubList returns a view of the portion of this list between the specified fromIndex, inclusive, and toIndex,
// exclusive.
func (s *SliceList[T]) ToSubList(fromIndex, toIndex int) List[T] {
	if fromIndex < 0 || toIndex >= len(s.Source) || toIndex < fromIndex {
		return nil
	}
	return &SliceList[T]{
		Source: s.Source[fromIndex : toIndex+1],
	}
}
