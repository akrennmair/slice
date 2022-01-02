package channel

import "context"

// MapWithContext does the same as Map, and receives a context.Context to be cancellable.
func MapWithContext[T1, T2 any](ctx context.Context, input <-chan T1, f func(T1) T2) <-chan T2 {
	ch := make(chan T2, 1)

	go func() {
		defer close(ch)
		for {
			select {
			case <-ctx.Done():
				return
			case elem, ok := <-input:
				if !ok {
					return
				}
				select {
				case <-ctx.Done():
					return
				case ch <- f(elem):
				}
			}
		}
	}()

	return ch
}

// Map returns a chan that produces the results of calling the provided function
// on every element read from the provided input chan.
func Map[T1, T2 any](input <-chan T1, f func(T1) T2) <-chan T2 {
	return MapWithContext(context.Background(), input, f)
}
