package channel

import "context"

func ToSliceWithContext[T any](ctx context.Context, inputC <-chan T) ([]T, error) {
	return ReduceWithContext(ctx, inputC, func(elems []T, v T) []T {
		return append(elems, v)
	})
}

func ToSlice[T any](in <-chan T) ([]T, error) {
	return ToSliceWithContext(context.Background(), in)
}
