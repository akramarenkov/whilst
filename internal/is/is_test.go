package is

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSpace(t *testing.T) {
	require.True(t, Space('\t'))
	require.True(t, Space('\n'))
	require.True(t, Space('\v'))
	require.True(t, Space('\f'))
	require.True(t, Space('\r'))
	require.True(t, Space(' '))
	require.True(t, Space(symbolNBSP))
	require.True(t, Space(symbolNEL))

	require.False(t, Space(0x8))
	require.False(t, Space(0xe))
	require.False(t, Space(0x1f))
	require.False(t, Space(0x21))
}

func TestDigit(t *testing.T) {
	require.True(t, Digit('0'))
	require.True(t, Digit('1'))
	require.True(t, Digit('2'))
	require.True(t, Digit('3'))
	require.True(t, Digit('4'))
	require.True(t, Digit('5'))
	require.True(t, Digit('6'))
	require.True(t, Digit('7'))
	require.True(t, Digit('8'))
	require.True(t, Digit('9'))

	require.False(t, Space('/'))
	require.False(t, Space(':'))
}
