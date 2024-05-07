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

func NewSliceFromCollection[T any](src collection.Collection[T]) []T {
	if src.IsEmpty() {
		return nil
	}

	iter := src.NewIterator()
	buf := make([]T, 0, src.Len())
	for iter.HasNext() {
		buf = append(buf, iter.Next())
	}
	return buf
}

func (s *SliceList[T]) NewIterator() collection.Iterator[T] {
	return NewIterator[T](s)
}

func (s *SliceList[T]) growIfRequired(n int) {
	if cap(s.Source)-len(s.Source) >= n {
		return
	}
	s.Source = append(make([]T, 0, len(s.Source)+n), s.Source...)
}

func (s *SliceList[T]) Add(v T) bool {
	s.Source = append(s.Source, v)
	return true
}

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

// Clear Removes all the elements from this collection. Does not de-allocates Source.
func (s *SliceList[T]) Clear() {
	s.Source = s.Source[:0]
}

func (s *SliceList[T]) Len() int {
	return len(s.Source)
}

func (s *SliceList[T]) IsEmpty() bool {
	return len(s.Source) == 0
}

func (s *SliceList[T]) ToSlice() []T {
	return s.Source
}

func (s *SliceList[T]) AddAt(index int, v T) {
	if !isValidIndex(index, len(s.Source)) {
		return
	}

	s.Source = append(s.Source[:index], append(s.Source[index:], []T{v}...)...)
}

func (s *SliceList[T]) SetAt(index int, v T) T {
	if !isValidIndex(index, len(s.Source)) {
		var zeroVal T
		return zeroVal
	}

	tmpVal := s.Source[index]
	s.Source[index] = v
	return tmpVal
}

func (s *SliceList[T]) AddAllAt(index int, src collection.Collection[T]) bool {
	insertionIndex := index + 1
	if !isValidIndex(insertionIndex, len(s.Source)) {
		return false
	}
	s.growIfRequired(src.Len())
	newSlice := NewSliceFromCollection(src)
	s.Source = append(s.Source[:insertionIndex], append(newSlice, s.Source[insertionIndex:]...)...)
	return true
}

func (s *SliceList[T]) GetAt(index int) T {
	if !isValidIndex(index, len(s.Source)) {
		var zeroVal T
		return zeroVal
	}
	return s.Source[index]
}

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

func (s *SliceList[T]) AddSlice(items ...T) bool {
	if len(items) == 0 {
		return false
	}
	s.growIfRequired(len(items))
	s.Source = append(s.Source, items...)
	return true
}

func (s *SliceList[T]) ForEach(predicateFunc collection.IterablePredicateFunc[T]) {
	for _, item := range s.Source {
		willBreak := predicateFunc(item)
		if willBreak {
			break
		}
	}
}

func (s *SliceList[T]) ForEachWithIndex(predicateFunc collection.IterablePredicateBiFunc[int, T]) {
	for i, item := range s.Source {
		willBreak := predicateFunc(i, item)
		if willBreak {
			break
		}
	}
}

func (s *SliceList[T]) ToSubList(fromIndex, toIndex int) List[T] {
	if fromIndex < 0 || toIndex >= len(s.Source) || toIndex < fromIndex {
		return nil
	}
	return &SliceList[T]{
		Source: s.Source[fromIndex:toIndex],
	}
}
