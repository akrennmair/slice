package slice_test

import (
	"testing"

	"github.com/akrennmair/slice"
	"github.com/stretchr/testify/require"
)

func TestToChan(t *testing.T) {
	var testData []int
	for i := 1; i < 100; i++ {
		testData = append(testData, i+i/3)
	}

	ch := slice.ToChan(testData)

	var output []int
	for v := range ch {
		output = append(output, v)
	}

	require.Equal(t, testData, output)
}
