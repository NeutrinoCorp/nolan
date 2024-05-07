package maps

import (
	"github.com/neutrinocorp/nolan/collection"
	"github.com/neutrinocorp/nolan/collection/list"
)

// LinkedHashMap Hash table and linked list implementation of the Map interface, with predictable iteration order.
// This implementation differs from HashMap in that it maintains a doubly-linked list running through all of its entries.
// This linked list defines the iteration ordering, which is normally the order in which keys were inserted into the
// map (insertion-order). Note that insertion order is not affected if a key is re-inserted
// into the map. (A key k is reinserted into a map m if Put(k, v) is invoked when ContainsKey(k)
// would return true immediately prior to the invocation.)
type LinkedHashMap[K comparable, V any] struct {
	hashMap HashMap[K, V]
	list    *list.DoublyLinkedList[K]
}

var _ Map[string, int] = &LinkedHashMap[string, int]{}

func NewLinkedHashMap[K comparable, V any]() *LinkedHashMap[K, V] {
	return &LinkedHashMap[K, V]{
		hashMap: HashMap[K, V]{},
		list:    &list.DoublyLinkedList[K]{},
	}
}

func (v *LinkedHashMap[K, V]) Get(key K) V {
	return v.hashMap.Get(key)
}

func (v *LinkedHashMap[K, V]) GetWithFallback(key K, fallbackValue V) V {
	return v.hashMap.GetWithFallback(key, fallbackValue)
}

func (v *LinkedHashMap[K, V]) Put(key K, val V) bool {
	v.hashMap.Put(key, val)
	v.list.Add(key)
	return true
}

func (v *LinkedHashMap[K, V]) PutIfAbsent(key K, val V) bool {
	wasMod := v.hashMap.PutIfAbsent(key, val)
	if !wasMod {
		return false
	}
	v.list.Add(key)
	return true
}

func (v *LinkedHashMap[K, V]) PutAll(src Map[K, V]) {
	src.ForEach(func(key K, val V) bool {
		v.Put(key, val)
		return false
	})
}

func (v *LinkedHashMap[K, V]) PutAllEntries(entries ...Entry[K, V]) {
	for _, entry := range entries {
		v.Put(entry.Key, entry.Value)
	}
}

func (v *LinkedHashMap[K, V]) Remove(key K) V {
	val, found := v.hashMap[key]
	if !found {
		return val
	}

	delete(v.hashMap, key)
	v.list.RemoveAt(list.IndexOf[K](v.list, key))
	return val
}

func (v *LinkedHashMap[K, V]) Replace(key K, val V) bool {
	return v.hashMap.Replace(key, val)
}

func (v *LinkedHashMap[K, V]) ContainsKey(key K) bool {
	return v.hashMap.ContainsKey(key)
}

func (v *LinkedHashMap[K, V]) Len() int {
	return v.hashMap.Len()
}

func (v *LinkedHashMap[K, V]) Clear() {
	v.hashMap.Clear()
	v.list.Clear()
}

func (v *LinkedHashMap[K, V]) Keys() collection.Collection[K] {
	return v.hashMap.Keys()
}

func (v *LinkedHashMap[K, V]) Values() collection.Collection[V] {
	return v.hashMap.Values()
}

func (v *LinkedHashMap[K, V]) ForEach(predicateFunc collection.IterablePredicateBiFunc[K, V]) {
	v.list.ForEach(func(mapKey K) bool {
		val := v.hashMap[mapKey]
		return predicateFunc(mapKey, val)
	})
}
