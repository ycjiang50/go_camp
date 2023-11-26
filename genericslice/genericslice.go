package genericslice

func Delete[T any](idx int, values []T) []T {
	if idx < 0 || idx >= len(values) {
		panic("Invalid index")
	}
	values = append(values[:idx], values[idx+1:]...)
	if len(values) <= cap(values)/2 {
		newValues := make([]T, len(values), len(values))
		for i, v := range values {
			newValues[i] = v
		}
		return newValues
	}
	return values
}
