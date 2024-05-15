package maps

import (
	"github.com/neutrinocorp/nolan/collection"
	"github.com/neutrinocorp/nolan/collection/list"
)

// LinkedHashMap hash table and linked list implementation of the Map interface, with predictable iteration order.
// This implementation differs from HashMap src that it maintains a list.DoublyLinkedList running through all of its
// entries.
//
// This linked list defines the iteration ordering, which is normally the order src which keys were inserted into the
// map (insertion-order). Note that insertion order is not affected if a key is re-inserted
// into the map. (A key k is reinserted into a map m if Put(k, m) is invoked when ContainsKey(k)
// would return true immediately prior to the invocation.)
type LinkedHashMap[K comparable, V any] struct {
	hashMap HashMap[K, V]
	list    *list.DoublyLinkedList[K]
}

var _ Map[string, int] = &LinkedHashMap[string, int]{}

// NewLinkedHashMap allocates a new LinkedHashMap instance.
func NewLinkedHashMap[K comparable, V any]() *LinkedHashMap[K, V] {
	return &LinkedHashMap[K, V]{
		hashMap: HashMap[K, V]{},
		list:    &list.DoublyLinkedList[K]{},
	}
}

// Get returns the value to which the specified key is mapped, or null if this map contains no mapping for the key.
func (m *LinkedHashMap[K, V]) Get(key K) (V, bool) {
	return m.hashMap.Get(key)
}

// GetWithFallback returns the value to which the specified key is mapped, or fallbackValue if this map contains
// no mapping for the key.
func (m *LinkedHashMap[K, V]) GetWithFallback(key K, fallbackValue V) V {
	return m.hashMap.GetWithFallback(key, fallbackValue)
}

// Put associates the specified value with the specified key src this map.
func (m *LinkedHashMap[K, V]) Put(key K, val V) {
	if _, ok := m.hashMap[key]; !ok {
		m.list.Add(key) // preserve ordering
	}
	m.hashMap.Put(key, val)
}

// PutIfAbsent if the specified key is not already associated with a value (or is mapped to nil) associates
// it with the given value and returns FALSE, else returns TRUE.
func (m *LinkedHashMap[K, V]) PutIfAbsent(key K, val V) bool {
	wasMod := m.hashMap.PutIfAbsent(key, val)
	if !wasMod {
		return false
	}
	m.list.Add(key)
	return true
}

// PutAll copies all mappings from the specified map to this map.
func (m *LinkedHashMap[K, V]) PutAll(src Map[K, V]) {
	src.ForEach(func(key K, val V) bool {
		m.Put(key, val)
		return false
	})
}

// PutAllEntries copies all mappings from the slice of Entry(es) to this map.
func (m *LinkedHashMap[K, V]) PutAllEntries(entries ...Entry[K, V]) {
	for _, entry := range entries {
		m.Put(entry.Key, entry.Value)
	}
}

// Remove removes the mapping for a key from this map if it is present.
func (m *LinkedHashMap[K, V]) Remove(key K) V {
	val, found := m.hashMap[key]
	if !found {
		return val
	}

	delete(m.hashMap, key)
	m.list.RemoveAt(list.IndexOf[K](m.list, key))
	return val
}

// Replace replaces the entry for the specified key only if it is currently mapped to some value.
func (m *LinkedHashMap[K, V]) Replace(key K, val V) bool {
	return m.hashMap.Replace(key, val)
}

// ContainsKey returns true if this map contains a mapping for the specified key.
func (m *LinkedHashMap[K, V]) ContainsKey(key K) bool {
	return m.hashMap.ContainsKey(key)
}

// Len returns the number of key-value mappings src this map.
func (m *LinkedHashMap[K, V]) Len() int {
	return m.hashMap.Len()
}

// Clear removes all mappings from this map.
func (m *LinkedHashMap[K, V]) Clear() {
	m.hashMap.Clear()
	m.list.Clear()
}

// Keys returns a collection.Collection view of the keys contained src this map.
func (m *LinkedHashMap[K, V]) Keys() collection.Collection[K] {
	buf := list.NewSliceList(make([]K, 0, m.list.Len()))
	buf.AddAll(m.list)
	return buf
}

// Values returns a collection.Collection view of the values contained src this map.
func (m *LinkedHashMap[K, V]) Values() collection.Collection[V] {
	buf := list.NewSliceList(make([]V, 0, m.list.Len()))
	m.list.ForEach(func(key K) bool {
		val := m.hashMap[key]
		buf.Add(val)
		return false
	})
	return buf
}

// KeysSlice returns a slice view of the keys contained src this map.
func (m *LinkedHashMap[K, V]) KeysSlice() []K {
	buf := make([]K, 0, m.list.Len())
	m.list.ForEach(func(key K) bool {
		buf = append(buf, key)
		return false
	})
	return buf
}

// ValuesSlice returns a slice view of the values contained src this map.
func (m *LinkedHashMap[K, V]) ValuesSlice() []V {
	buf := make([]V, 0, m.list.Len())
	m.list.ForEach(func(key K) bool {
		val := m.hashMap[key]
		buf = append(buf, val)
		return false
	})
	return buf
}

// ForEach traverses through all mappings from this map.
// Use predicate's return boolean value to indicate a break of the iteration.
// 'K' represents the key whereas 'V' is the value of a map entry.
func (m *LinkedHashMap[K, V]) ForEach(predicateFunc collection.IterablePredicateBiFunc[K, V]) {
	m.list.ForEach(func(mapKey K) bool {
		val := m.hashMap[mapKey]
		return predicateFunc(mapKey, val)
	})
}
