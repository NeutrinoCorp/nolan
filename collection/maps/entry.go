package maps

// Entry a Map item.
//
// K: Key's type.
//
// V: Value's type.
type Entry[K comparable, V any] struct {
	Key   K
	Value V
}
