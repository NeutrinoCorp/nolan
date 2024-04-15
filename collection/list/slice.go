package list

import "github.com/neutrinocorp/nolan/collection"

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

func (s *SliceList[T]) Add(v T) bool {
	s.Source = append(s.Source, v)
	return true
}

func (s *SliceList[T]) AddAll(src collection.Collection[T]) bool {
	iter := src.NewIterator()
	wasMod := iter.HasNext()
	for iter.HasNext() {
		s.Source = append(s.Source, iter.Next())
	}
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
	if index > len(s.Source)-1 || index < 0 {
		return
	}

	s.Source = append(s.Source[:index], append(s.Source[index:], []T{v}...)...)
}

func (s *SliceList[T]) SetAt(index int, v T) T {
	if index > len(s.Source)-1 || index < 0 {
		var zeroVal T
		return zeroVal
	}

	tmpVal := s.Source[index]
	s.Source[index] = v
	return tmpVal
}

func (s *SliceList[T]) growIfRequired(n int) {
	if cap(s.Source)-len(s.Source) >= n {
		return
	}
	s.Source = append(make([]T, 0, len(s.Source)+n), s.Source...)
}

func (s *SliceList[T]) AddAllAt(index int, src collection.Collection[T]) bool {
	insertionIndex := index + 1
	if insertionIndex > len(s.Source)-1 || insertionIndex < 0 {
		return false
	}
	s.growIfRequired(src.Len())

	newSlice := NewSliceFromCollection(src)
	s.Source = append(s.Source[:insertionIndex], append(newSlice, s.Source[insertionIndex:]...)...)
	return true
}

func (s *SliceList[T]) GetAt(index int) T {
	if index > len(s.Source)-1 || index < 0 {
		var zeroVal T
		return zeroVal
	}

	return s.Source[index]
}

func (s *SliceList[T]) RemoveAt(index int) T {
	item := s.GetAt(index)
	copy(s.Source[index:], s.Source[index+1:])
	var zeroVal T
	s.Source[len(s.Source)-1] = zeroVal // or the zero value of T
	s.Source = s.Source[:len(s.Source)-1]
	return item
}

func (s *SliceList[T]) ToSubList(fromIndex, toIndex int) List[T] {
	// TODO: IMPLEMENT ME!
	return nil
}
