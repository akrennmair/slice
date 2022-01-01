package slice

import "context"

// ToChanWithContext does the same as ToChan, but is cancellable through
// the provided context.Context object.
func ToChanWithContext[T any](ctx context.Context, in []T) <-chan T {
	ch := make(chan T, 1)

	go func() {
		defer close(ch)
		for _, v := range in {
			select {
			case <-ctx.Done():
				return
			case ch <- v:
			}
		}
	}()

	return ch
}

// ToChan returns a chan that produces all elements of the provided
// input slice in order.
func ToChan[T any](in []T) <-chan T {
	return ToChanWithContext(context.Background(), in)
}
