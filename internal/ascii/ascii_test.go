package ascii

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsSpace(t *testing.T) {
	require.True(t, IsSpace('\t'))
	require.True(t, IsSpace('\n'))
	require.True(t, IsSpace('\v'))
	require.True(t, IsSpace('\f'))
	require.True(t, IsSpace('\r'))
	require.True(t, IsSpace(' '))

	require.False(t, IsSpace(0x8))
	require.False(t, IsSpace(0xe))
	require.False(t, IsSpace(0x1f))
	require.False(t, IsSpace(0x21))
}

func TestIsDigit(t *testing.T) {
	require.True(t, IsDigit('0'))
	require.True(t, IsDigit('1'))
	require.True(t, IsDigit('2'))
	require.True(t, IsDigit('3'))
	require.True(t, IsDigit('4'))
	require.True(t, IsDigit('5'))
	require.True(t, IsDigit('6'))
	require.True(t, IsDigit('7'))
	require.True(t, IsDigit('8'))
	require.True(t, IsDigit('9'))

	require.False(t, IsSpace('/'))
	require.False(t, IsSpace(':'))
}

func TestByteToDigit(t *testing.T) {
	require.Equal(t, int8(0), ByteToDigit[int8]('0'))
	require.Equal(t, int8(1), ByteToDigit[int8]('1'))
	require.Equal(t, int8(2), ByteToDigit[int8]('2'))
	require.Equal(t, int8(3), ByteToDigit[int8]('3'))
	require.Equal(t, int8(4), ByteToDigit[int8]('4'))
	require.Equal(t, int8(5), ByteToDigit[int8]('5'))
	require.Equal(t, int8(6), ByteToDigit[int8]('6'))
	require.Equal(t, int8(7), ByteToDigit[int8]('7'))
	require.Equal(t, int8(8), ByteToDigit[int8]('8'))
	require.Equal(t, int8(9), ByteToDigit[int8]('9'))

	require.Equal(t, uint8(0), ByteToDigit[uint8]('0'))
	require.Equal(t, uint8(1), ByteToDigit[uint8]('1'))
	require.Equal(t, uint8(2), ByteToDigit[uint8]('2'))
	require.Equal(t, uint8(3), ByteToDigit[uint8]('3'))
	require.Equal(t, uint8(4), ByteToDigit[uint8]('4'))
	require.Equal(t, uint8(5), ByteToDigit[uint8]('5'))
	require.Equal(t, uint8(6), ByteToDigit[uint8]('6'))
	require.Equal(t, uint8(7), ByteToDigit[uint8]('7'))
	require.Equal(t, uint8(8), ByteToDigit[uint8]('8'))
	require.Equal(t, uint8(9), ByteToDigit[uint8]('9'))
}

func TestDigitToByte(t *testing.T) {
	require.Equal(t, byte('0'), DigitToByte(int8(0)))
	require.Equal(t, byte('1'), DigitToByte(int8(1)))
	require.Equal(t, byte('2'), DigitToByte(int8(2)))
	require.Equal(t, byte('3'), DigitToByte(int8(3)))
	require.Equal(t, byte('4'), DigitToByte(int8(4)))
	require.Equal(t, byte('5'), DigitToByte(int8(5)))
	require.Equal(t, byte('6'), DigitToByte(int8(6)))
	require.Equal(t, byte('7'), DigitToByte(int8(7)))
	require.Equal(t, byte('8'), DigitToByte(int8(8)))
	require.Equal(t, byte('9'), DigitToByte(int8(9)))

	require.Equal(t, byte('0'), DigitToByte(uint8(0)))
	require.Equal(t, byte('1'), DigitToByte(uint8(1)))
	require.Equal(t, byte('2'), DigitToByte(uint8(2)))
	require.Equal(t, byte('3'), DigitToByte(uint8(3)))
	require.Equal(t, byte('4'), DigitToByte(uint8(4)))
	require.Equal(t, byte('5'), DigitToByte(uint8(5)))
	require.Equal(t, byte('6'), DigitToByte(uint8(6)))
	require.Equal(t, byte('7'), DigitToByte(uint8(7)))
	require.Equal(t, byte('8'), DigitToByte(uint8(8)))
	require.Equal(t, byte('9'), DigitToByte(uint8(9)))
}
