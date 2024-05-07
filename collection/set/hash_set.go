package set

import (
	"maps"

	"github.com/neutrinocorp/nolan/collection"
)

type HashSet[K comparable] map[K]struct{}

var _ Set[string] = HashSet[string]{}

func (h HashSet[K]) NewIterator() collection.Iterator[K] {
	return NewIterator[K](h)
}

func (h HashSet[K]) Add(v K) bool {
	if _, ok := h[v]; ok {
		return false
	}
	h[v] = struct{}{}
	return true
}

func (h HashSet[K]) AddAll(src collection.Collection[K]) bool {
	wasMod := false
	src.ForEach(func(a K) bool {
		wasModItem := h.Add(a)
		if wasModItem {
			wasMod = true
		}
		return false
	})
	return wasMod
}

func (h HashSet[K]) AddSlice(items ...K) bool {
	wasMod := false
	for _, item := range items {
		wasModItem := h.Add(item)
		if wasModItem {
			wasMod = true
		}
	}
	return wasMod
}

func (h HashSet[K]) Clear() {
	maps.DeleteFunc(h, func(_ K, _ struct{}) bool {
		return true
	})
}

func (h HashSet[K]) Len() int {
	return len(h)
}

func (h HashSet[K]) IsEmpty() bool {
	return len(h) == 0
}

func (h HashSet[K]) ToSlice() []K {
	buf := make([]K, 0, len(h))
	for k := range h {
		buf = append(buf, k)
	}
	return buf
}

func (h HashSet[K]) Contains(v K) bool {
	_, ok := h[v]
	return ok
}

func (h HashSet[K]) ContainsAll(src collection.Collection[K]) bool {
	iter := src.NewIterator()
	for iter.HasNext() {
		key := iter.Next()
		if _, ok := h[key]; !ok {
			return false
		}
	}
	return true
}

func (h HashSet[K]) ContainsSlice(src ...K) bool {
	for _, item := range src {
		if _, ok := h[item]; !ok {
			return false
		}
	}
	return true
}

func (h HashSet[K]) ForEach(predicateFunc collection.IterablePredicateFunc[K]) {
	for k := range h {
		willBreak := predicateFunc(k)
		if willBreak {
			break
		}
	}
}
