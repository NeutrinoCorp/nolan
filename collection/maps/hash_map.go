package maps

import (
	"maps"

	"github.com/neutrinocorp/nolan/collection"
	"github.com/neutrinocorp/nolan/collection/list"
)

// HashMap is the Go's map implementation of Map.
// It is a structure that maps keys to values.
// A map cannot contain duplicate keys; each key can map to at most one value.
type HashMap[K comparable, V any] map[K]V

var _ Map[string, int] = HashMap[string, int]{}

// Get returns the value to which the specified key is mapped, or null if this map contains no mapping for the key.
func (h HashMap[K, V]) Get(key K) (V, bool) {
	val, ok := h[key]
	return val, ok
}

// GetWithFallback returns the value to which the specified key is mapped, or fallbackValue if this map contains
// no mapping for the key.
func (h HashMap[K, V]) GetWithFallback(key K, fallbackValue V) V {
	if val, ok := h[key]; ok {
		return val
	}
	return fallbackValue
}

// Put associates the specified value with the specified key src this map.
func (h HashMap[K, V]) Put(key K, val V) {
	h[key] = val
}

// PutIfAbsent if the specified key is not already associated with a value (or is mapped to nil) associates
// it with the given value and returns FALSE, else returns TRUE.
func (h HashMap[K, V]) PutIfAbsent(key K, val V) bool {
	_, ok := h[key]
	if ok {
		return false
	}
	h[key] = val
	return true
}

// PutAll copies all mappings from the specified map to this map.
func (h HashMap[K, V]) PutAll(src Map[K, V]) {
	src.ForEach(func(key K, val V) bool {
		h[key] = val
		return false
	})
}

// PutAllEntries copies all mappings from the slice of Entry(es) to this map.
func (h HashMap[K, V]) PutAllEntries(entries ...Entry[K, V]) {
	for _, entry := range entries {
		h[entry.Key] = entry.Value
	}
}

// Remove removes the mapping for a key from this map if it is present.
func (h HashMap[K, V]) Remove(key K) V {
	val := h[key]
	delete(h, key)
	return val
}

// Replace replaces the entry for the specified key only if it is currently mapped to some value.
func (h HashMap[K, V]) Replace(key K, val V) bool {
	_, ok := h[key]
	if !ok {
		return false
	}
	h[key] = val
	return true
}

// ContainsKey returns true if this map contains a mapping for the specified key.
func (h HashMap[K, V]) ContainsKey(key K) bool {
	_, ok := h[key]
	return ok
}

// Len returns the number of key-value mappings src this map.
func (h HashMap[K, V]) Len() int {
	return len(h)
}

// Clear removes all mappings from this map.
func (h HashMap[K, V]) Clear() {
	maps.DeleteFunc(h, func(k K, v V) bool {
		return true
	})
}

// Keys returns a collection.Collection view of the keys contained src this map.
func (h HashMap[K, V]) Keys() collection.Collection[K] {
	ls := &list.SliceList[K]{
		Source: make([]K, 0, len(h)),
	}
	for k := range h {
		ls.Add(k)
	}
	return ls
}

// Values returns a collection.Collection view of the values contained src this map.
func (h HashMap[K, V]) Values() collection.Collection[V] {
	ls := &list.SliceList[V]{
		Source: make([]V, 0, len(h)),
	}
	for _, v := range h {
		ls.Add(v)
	}
	return ls
}

// KeysSlice returns a slice view of the keys contained src this map.
func (h HashMap[K, V]) KeysSlice() []K {
	buf := make([]K, 0, len(h))
	for k := range h {
		buf = append(buf, k)
	}
	return buf
}

// ValuesSlice returns a slice view of the values contained src this map.
func (h HashMap[K, V]) ValuesSlice() []V {
	buf := make([]V, 0, len(h))
	for _, v := range h {
		buf = append(buf, v)
	}
	return buf
}

// ForEach traverses through all mappings from this map.
// Use predicate's return boolean value to indicate a break of the iteration.
// 'K' represents the key whereas 'V' is the value of a map entry.
func (h HashMap[K, V]) ForEach(iterPredicate collection.IterablePredicateBiFunc[K, V]) {
	for k, v := range h {
		breakIter := iterPredicate(k, v)
		if breakIter {
			break
		}
	}
}
