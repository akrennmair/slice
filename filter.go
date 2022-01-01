package slice

func Filter[T any](input []T, pred func(T) bool) (output []T) {
	for _, v := range input {
		if pred(v) {
			output = append(output, v)
		}
	}
	return output
}
