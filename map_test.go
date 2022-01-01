package slice_test

import (
	"context"
	"testing"

	"github.com/akrennmair/slice"
	"github.com/stretchr/testify/require"
)

func TestMap(t *testing.T) {
	input, expectedResult := []int32{2, 4, 8, 16}, []int64{2, 4, 8, 16}
	result := slice.Map(input, func(v int32) int64 {
		return int64(v)
	})

	require.Equal(t, expectedResult, result)
}

func TestMapConcurrent(t *testing.T) {
	input, expectedResult := []int32{2, 4, 8, 16}, []int64{2, 4, 8, 16}
	result := slice.MapConcurrent(input, func(v int32) int64 {
		return int64(v)
	})

	require.Equal(t, expectedResult, result)
}

func TestMapConcurrentWithContext(t *testing.T) {
	input, expectedResult := []int32{2, 4, 8, 16}, []int64{2, 4, 8, 16}
	result := slice.MapConcurrentWithContext(context.Background(), input, func(v int32) int64 {
		return int64(v)
	})

	require.Equal(t, expectedResult, result)
}

func TestMapConcurrentWithCancelledContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	input, expectedResult := []int32{2, 4, 8, 16}, []int64(nil)
	result := slice.MapConcurrentWithContext(ctx, input, func(v int32) int64 {
		return int64(v)
	})

	require.Equal(t, expectedResult, result)
}

func BenchmarkMap(b *testing.B) {
	b.StopTimer()
	input := make([]int64, b.N)
	b.StartTimer()
	slice.Map(input, func(v int64) int64 {
		return v
	})
}

func BenchmarkMapConcurrentWithContext(b *testing.B) {
	b.StopTimer()
	input := make([]int64, b.N)
	b.StartTimer()
	slice.MapConcurrentWithContext(context.Background(), input, func(v int64) int64 {
		return v
	})
}

func BenchmarkMapConcurrent(b *testing.B) {
	b.StopTimer()
	input := make([]int64, b.N)
	b.StartTimer()
	slice.MapConcurrent(input, func(v int64) int64 {
		return v
	})
}
