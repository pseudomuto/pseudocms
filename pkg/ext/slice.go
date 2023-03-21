package ext

// MapSlice maps the given slice from type T to type U and returns a new slice.
func MapSlice[T, U any](sl []T, fn func(T) U) []U {
	res, _ := MapSliceErr(sl, func(t T) (U, error) { return fn(t), nil })
	return res
}

// MapSliceErr maps the given slice from type T to type U and returns a new slice.
// If fn returns an error, it will be returned immediately and the mapping
// operation will be stopped.
func MapSliceErr[T, U any](sl []T, fn func(T) (U, error)) ([]U, error) {
	newSl := make([]U, len(sl))
	for i, v := range sl {
		res, err := fn(v)
		if err != nil {
			return nil, err
		}

		newSl[i] = res
	}

	return newSl, nil
}
