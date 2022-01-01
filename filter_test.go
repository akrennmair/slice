package slice_test

import (
	"testing"

	"github.com/akrennmair/slice"
	"github.com/stretchr/testify/require"
)

func TestFilter(t *testing.T) {
	var testData []int
	for i := 1; i < 100; i++ {
		testData = append(testData, i)
	}

	output := slice.Filter(testData, func(v int) bool {
		return v >= 30 && v <= 33
	})

	require.Equal(t, []int{30, 31, 32, 33}, output)
}
