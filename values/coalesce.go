package values

// Coalesce returns the first non-nil (or non-zero) value from a given list.
func Coalesce[T comparable](vals ...T) T {
	var zeroVal T
	for _, v := range vals {
		if v != zeroVal {
			return v
		}
	}
	return zeroVal
}
