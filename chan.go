package slice

import "context"

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

func ToChan[T any](in []T) <-chan T {
	return ToChanWithContext(context.Background(), in)
}
