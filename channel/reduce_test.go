package channel_test

import (
	"testing"

	"github.com/akrennmair/slice"
	"github.com/akrennmair/slice/channel"
	"github.com/stretchr/testify/require"
)

func TestReduce(t *testing.T) {
	var testData []int
	for i := 1; i <= 100; i++ {
		testData = append(testData, i)
	}

	sum, err := channel.Reduce(slice.ToChan(testData), func(acc int, v int) int {
		return acc + v
	})
	require.NoError(t, err)
	require.Equal(t, 5050, sum)
}
