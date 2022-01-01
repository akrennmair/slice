package channel

import "context"

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

func Reduce[T1, T2 any](inputC <-chan T1, f func(T2, T1) T2) (T2, error) {
	return ReduceWithContext(context.Background(), inputC, f)
}
