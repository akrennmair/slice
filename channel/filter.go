package channel

import "context"

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

func Filter[T any](inputC <-chan T, pred func(T) bool) <-chan T {
	return FilterWithContext(context.Background(), inputC, pred)
}
