package channel

import "context"

// ReduceWithContext does the same as Reduce, and is cancellable with a context.Context object.
func ReduceWithContext[T1, T2 any](ctx context.Context, inputC <-chan T1, f func(T2, T1) T2) (T2, error) {
	var acc T2

loop:
	for {
		select {
		case v, ok := <-inputC:
			if !ok {
				break loop
			}
			acc = f(acc, v)
		case <-ctx.Done():
			return acc, ctx.Err()
		}
	}

	return acc, nil
}

// Reduce executes a provided function on each element read from the
// provided chan in order,passing the return value of the previous
// function call on the preceding element. The final result of running the
// provided function across all elements is returned.
func Reduce[T1, T2 any](inputC <-chan T1, f func(T2, T1) T2) (T2, error) {
	return ReduceWithContext(context.Background(), inputC, f)
}
