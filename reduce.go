package slice

func Reduce[T1, T2 any](input []T1, f func(T2, T1) T2) T2 {
	var acc T2

	for _, v := range input {
		acc = f(acc, v)
	}

	return acc
}
