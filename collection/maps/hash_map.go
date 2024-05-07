package maps

import (
	"maps"

	"github.com/neutrinocorp/nolan/collection"
	"github.com/neutrinocorp/nolan/collection/list"
)

type HashMap[K comparable, V any] map[K]V

var _ Map[string, int] = HashMap[string, int]{}

func (h HashMap[K, V]) Get(key K) V {
	val, _ := h[key]
	return val
}

func (h HashMap[K, V]) GetWithFallback(key K, fallbackValue V) V {
	if val, ok := h[key]; ok {
		return val
	}
	return fallbackValue
}

func (h HashMap[K, V]) Put(key K, val V) bool {
	h[key] = val
	return false
}

func (h HashMap[K, V]) PutIfAbsent(key K, val V) bool {
	_, ok := h[key]
	if ok {
		return false
	}
	h[key] = val
	return true
}

func (h HashMap[K, V]) PutAll(src Map[K, V]) {
	src.ForEach(func(key K, val V) bool {
		h[key] = val
		return false
	})
}

func (h HashMap[K, V]) PutAllEntries(entries ...Entry[K, V]) {
	for _, entry := range entries {
		h[entry.Key] = entry.Value
	}
}

func (h HashMap[K, V]) Remove(key K) V {
	val, _ := h[key]
	delete(h, key)
	return val
}

func (h HashMap[K, V]) Replace(key K, val V) bool {
	_, ok := h[key]
	if !ok {
		return false
	}
	h[key] = val
	return true
}

func (h HashMap[K, V]) ContainsKey(key K) bool {
	_, ok := h[key]
	return ok
}

func (h HashMap[K, V]) Len() int {
	return len(h)
}

func (h HashMap[K, V]) Clear() {
	maps.DeleteFunc(h, func(k K, v V) bool {
		return true
	})
}

func (h HashMap[K, V]) Keys() collection.Collection[K] {
	ls := &list.SliceList[K]{
		Source: make([]K, 0, len(h)),
	}
	for k := range h {
		ls.Add(k)
	}
	return ls
}

func (h HashMap[K, V]) Values() collection.Collection[V] {
	ls := &list.SliceList[V]{
		Source: make([]V, 0, len(h)),
	}
	for _, v := range h {
		ls.Add(v)
	}
	return ls
}

func (h HashMap[K, V]) ForEach(iterPredicate collection.IterablePredicateBiFunc[K, V]) {
	for k, v := range h {
		breakIter := iterPredicate(k, v)
		if breakIter {
			break
		}
	}
}
