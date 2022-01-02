package channel

import "context"

// FilterWithContext does the same as Filter, and is cancellable with a context.Context object.
func FilterWithContext[T any](ctx context.Context, inputC <-chan T, pred func(T) bool) <-chan T {
	outputC := make(chan T, 1)

	go func() {
		defer close(outputC)

		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-inputC:
				if !ok {
					return
				}
				if pred(v) {
					select {
					case <-ctx.Done():
						return
					case outputC <- v:
					}
				}
			}
		}
	}()

	return outputC
}

// Filter returns a chan that produces all elements from the
// input chan for which the provided predicate function returns true.
func Filter[T any](inputC <-chan T, pred func(T) bool) <-chan T {
	return FilterWithContext(context.Background(), inputC, pred)
}
