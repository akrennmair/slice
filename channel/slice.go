package channel

import "context"

// ToSliceWithContext does the same as ToSlice, and is cancellable with a
// context.Context object.
func ToSliceWithContext[T any](ctx context.Context, inputC <-chan T) ([]T, error) {
	return ReduceWithContext(ctx, inputC, func(elems []T, v T) []T {
		return append(elems, v)
	})
}

// ToSlice reads all elements from the provided input chan and returns them as a slice.
func ToSlice[T any](in <-chan T) ([]T, error) {
	return ToSliceWithContext(context.Background(), in)
}
