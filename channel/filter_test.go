package channel_test

import (
	"testing"

	"github.com/akrennmair/slice"
	"github.com/akrennmair/slice/channel"
	"github.com/stretchr/testify/require"
)

func TestFilter(t *testing.T) {
	var testData []int
	for i := 1; i <= 100; i++ {
		testData = append(testData, i)
	}

	outputData, err := channel.ToSlice(channel.Filter(slice.ToChan(testData), func(v int) bool {
		return v >= 42 && v <= 45
	}))
	require.NoError(t, err)
	require.Equal(t, []int{42, 43, 44, 45}, outputData)
}
