package convert

import "github.com/neutrinocorp/nolan/function"

// ConverterFunc is a function.DelegateFunc type used to convert values from A type to B types.
type ConverterFunc[A, B any] function.DelegateFunc[A, B]

// ConverterSafeFunc is a function.DelegateSafeFunc type used to convert values from A type to B types. Returns
// error if something fails.
type ConverterSafeFunc[A, B any] function.DelegateSafeFunc[A, B]

// ConvertMany converts src list to a B type list.
func ConvertMany[A, B any](src []A, convertFunc ConverterFunc[A, B]) []B {
	if len(src) == 0 {
		return nil
	} else if convertFunc == nil {
		panic("convertFunc cannot be nil")
	}

	dst := make([]B, 0, len(src))
	for _, item := range src {
		convertedItem := convertFunc(item)
		dst = append(dst, convertedItem)
	}
	return dst
}

// ConvertManySafe converts src list to a B type list. Returns error if any of the conversions fail.
func ConvertManySafe[A, B any](src []A, convertFunc ConverterSafeFunc[A, B]) ([]B, error) {
	if len(src) == 0 {
		return nil, nil
	} else if convertFunc == nil {
		panic("convertFunc cannot be nil")
	}

	dst := make([]B, 0, len(src))
	for _, item := range src {
		convertedItem, err := convertFunc(item)
		if err != nil {
			return nil, err
		}
		dst = append(dst, convertedItem)
	}
	return dst, nil
}
