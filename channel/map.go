package channel

import "context"

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

func Map[T1, T2 any](input <-chan T1, f func(T1) T2) <-chan T2 {
	return MapWithContext(context.Background(), input, f)
}
