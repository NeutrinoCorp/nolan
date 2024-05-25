package values

// NewPtr returns a pointer to src.
func NewPtr[T any](src T) *T {
	return &src
}

// NewNillablePtr returns a pointer to src. If src is zero-value, then nil is returned.
func NewNillablePtr[T comparable](src T) *T {
	var zeroVal T
	if src == zeroVal {
		return nil
	}

	return &src
}
