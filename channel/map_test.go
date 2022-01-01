package channel_test

import (
	"testing"

	"github.com/akrennmair/slice"
	"github.com/akrennmair/slice/channel"
	"github.com/stretchr/testify/require"
)

func TestMap(t *testing.T) {
	input, expectedOutput := []int{23, 42, 9001}, []int64{46, 84, 18002}
	output, err := channel.ToSlice(channel.Map(slice.ToChan(input), func(v int) int64 {
		return int64(v * 2)
	}))
	require.NoError(t, err)
	require.Equal(t, expectedOutput, output)
}
