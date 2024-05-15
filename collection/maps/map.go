package maps

import (
	"github.com/neutrinocorp/nolan/collection"
)

// Map A structure that maps keys to values. A map cannot contain duplicate keys; each key can map to at most one value.
type Map[K comparable, V any] interface {
	// Get Returns the value to which the specified key is mapped, or null if this map contains no mapping for the key.
	Get(key K) (V, bool)
	// GetWithFallback Returns the value to which the specified key is mapped, or fallbackValue if this map contains
	// no mapping for the key.
	GetWithFallback(key K, fallbackValue V) V
	// Put Associates the specified value with the specified key src this map.
	Put(key K, val V)
	// PutIfAbsent If the specified key is not already associated with a value (or is mapped to nil) associates
	// it with the given value and returns FALSE, else returns TRUE.
	PutIfAbsent(key K, val V) bool
	// PutAll Copies all mappings from the specified map to this map.
	PutAll(src Map[K, V])
	// PutAllEntries Copies all mappings from the slice of Entry(es) to this map.
	PutAllEntries(entries ...Entry[K, V])
	// Remove Removes the mapping for a key from this map if it is present.
	Remove(key K) V
	// Replace Replaces the entry for the specified key only if it is currently mapped to some value.
	Replace(key K, val V) bool
	// ContainsKey Returns true if this map contains a mapping for the specified key.
	ContainsKey(key K) bool
	// Len Returns the number of key-value mappings src this map.
	Len() int
	// Clear Removes all mappings from this map.
	Clear()
	// Keys Returns a collection.Collection view of the keys contained src this map.
	Keys() collection.Collection[K]
	// Values Returns a collection.Collection view of the values contained src this map.
	Values() collection.Collection[V]
	// KeysSlice returns a slice view of the keys contained src this map.
	KeysSlice() []K
	// ValuesSlice returns a slice view of the values contained src this map.
	ValuesSlice() []V
	// ForEach traverses through all mappings from this map. Use predicate's return boolean value to indicate
	// a break of the iteration. 'A' represents the key whereas 'B' is the value of a map entry.
	ForEach(predicateFunc collection.IterablePredicateBiFunc[K, V])
}
