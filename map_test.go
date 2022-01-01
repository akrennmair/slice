package slice_test

import (
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

func BenchmarkMap(b *testing.B) {
	b.StopTimer()
	input := make([]int64, b.N)
	b.StartTimer()
	output := slice.Map(input, func(v int64) int64 {
		return v
	})
	_ = output
}

func BenchmarkMapConcurrent(b *testing.B) {
	b.StopTimer()
	input := make([]int64, b.N)
	b.StartTimer()
	output := slice.MapConcurrent(input, func(v int64) int64 {
		return v
	})
	_ = output
}
