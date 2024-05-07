package exception

type container interface {
	Unwrap() []error
}

// UnwrapErrors returns errors joined by errors.Join. If err was not joined by previous routine, then a slice
// with a single item will be returned containing solely err.
func UnwrapErrors(err error) []error {
	if err == nil {
		return nil
	}

	cont, ok := err.(container)
	if !ok {
		return []error{err}
	}

	return cont.Unwrap()
}
