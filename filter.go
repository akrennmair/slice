package slice

// Filter returns a new slice with all elements from the from the
// input elements for which the provided predicate function returns true.
func Filter[T any](input []T, pred func(T) bool) (output []T) {
	for _, v := range input {
		if pred(v) {
			output = append(output, v)
		}
	}
	return output
}
