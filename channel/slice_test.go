package channel_test

import (
	"testing"

	"github.com/akrennmair/slice"
	"github.com/akrennmair/slice/channel"
	"github.com/stretchr/testify/require"
)

func TestToSlice(t *testing.T) {
	input := []int{23, 42, 9001}
	ch := slice.ToChan(input)
	output, err := channel.ToSlice(ch)
	require.NoError(t, err)
	require.Equal(t, input, output)
}
