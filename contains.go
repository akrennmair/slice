package slice

// Contains returns a bool if input == X => true for any X in lst
func Contains[T comparable](val T, lst []T) bool {
	return Reduce(lst, func(contains bool, t T) bool {
		if contains {
			return true
		}
		return t == val
	})
}
