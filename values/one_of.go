package values

// IsOneOf returns if src is equals to any of the items contained in cmpValues variadic slice.
func IsOneOf[T comparable](src T, cmpValues ...T) bool {
	for _, comparison := range cmpValues {
		if src == comparison {
			return true
		}
	}
	return false
}
