package slice_test

import (
	"testing"

	"github.com/akrennmair/slice"
	"github.com/stretchr/testify/require"
)

func TestReduce(t *testing.T) {
	var testData []int8
	for i := int8(1); i < 100; i++ {
		testData = append(testData, i)
	}

	sum := slice.Reduce(testData, func(acc int16, v int8) int16 {
		return acc + int16(v)
	})

	require.Equal(t, int16(4950), sum)
}

func TestReduceWithInitialValue(t *testing.T) {
	var testData []int
	for i := 1; i <= 10; i++ {
		testData = append(testData, i)
	}

	result := slice.ReduceWithInitialValue(testData, 1, func(acc int, v int) int {
		return acc * v
	})

	require.Equal(t, 3628800, result)
}
