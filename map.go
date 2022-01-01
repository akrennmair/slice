package slice

import (
	"context"
	"sync"
)

// Map returns a new slice populated with the result of calling the provided function
// on every element in the provided input slice.
func Map[T1, T2 any](input []T1, f func(T1) T2) (output []T2) {
	for _, v := range input {
		output = append(output, f(v))
	}
	return output
}

// MapConcurrentWithContext does the same as Map, but concurrently, and receives a context.Context to be
// cancellable.
func MapConcurrentWithContext[T1, T2 any](ctx context.Context, input []T1, f func(T1) T2) (output []T2) {
	elemOrder := make(chan chan T2, len(input))

	go func() {
		defer close(elemOrder)

		var wg sync.WaitGroup

		for _, v := range input {
			elemC := make(chan T2, 1)
			wg.Add(1)
			go func(elemC chan<- T2, v T1) {
				defer wg.Done()
				select {
				case <-ctx.Done():
					return
				case elemC <- f(v):
				}
			}(elemC, v)
			select {
			case <-ctx.Done():
				return
			case elemOrder <- elemC:
			}
		}

		wg.Wait()
	}()

loop:
	for {
		select {
		case <-ctx.Done():
			break loop
		case elemC, ok := <-elemOrder:
			if !ok {
				break loop
			}
			select {
			case <-ctx.Done():
				break loop
			case elem := <-elemC:
				output = append(output, elem)
			}
		}
	}

	return output
}

// MapConcurrent does the same as Map, but concurrently.
func MapConcurrent[T1, T2 any](input []T1, f func(T1) T2) (output []T2) {
	elemOrder := make(chan chan T2, len(input))

	go func() {
		defer close(elemOrder)

		var wg sync.WaitGroup

		for _, v := range input {
			elemC := make(chan T2, 1)
			wg.Add(1)
			go func(elemC chan<- T2, v T1) {
				elemC <- f(v)
				wg.Done()
			}(elemC, v)
			elemOrder <- elemC
		}

		wg.Wait()
	}()

	for elemC := range elemOrder {
		output = append(output, <-elemC)
	}

	return output
}
