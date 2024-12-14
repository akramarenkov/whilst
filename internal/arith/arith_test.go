package arith

import (
	"math"
	"testing"
	"time"

	"github.com/akramarenkov/whilst/internal/consts"
	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	sum, err := Add(math.MaxInt64, 0)
	require.NoError(t, err)
	require.Equal(t, int64(math.MaxInt64), sum)

	sum, err = Add(math.MaxInt64, 1)
	require.Error(t, err)
	require.Equal(t, int64(0), sum)

	sum, err = Add(math.MaxInt64, math.MaxInt64)
	require.Error(t, err)
	require.Equal(t, int64(0), sum)
}

func TestAddDuration(t *testing.T) {
	sum, err := AddDuration(math.MaxInt64, 0)
	require.NoError(t, err)
	require.Equal(t, time.Duration(math.MaxInt64), sum)

	sum, err = AddDuration(math.MaxInt64, 1)
	require.Error(t, err)
	require.Equal(t, time.Duration(0), sum)

	sum, err = AddDuration(math.MaxInt64, math.MaxInt64)
	require.Error(t, err)
	require.Equal(t, time.Duration(0), sum)
}

func TestMulBy10(t *testing.T) {
	sum, err := MulBy10(math.MaxInt64 / consts.DecimalBase)
	require.NoError(t, err)
	require.Equal(t, int64(9223372036854775800), sum)

	sum, err = MulBy10(math.MaxInt64/consts.DecimalBase + 1)
	require.Error(t, err)
	require.Equal(t, int64(0), sum)
}

func TestMulByMicrosecond(t *testing.T) {
	sum, err := MulByMicrosecond(math.MaxInt64 / time.Microsecond)
	require.NoError(t, err)
	require.Equal(t, time.Duration(9223372036854775000), sum)

	sum, err = MulByMicrosecond(math.MaxInt64/time.Microsecond + 1)
	require.Error(t, err)
	require.Equal(t, time.Duration(0), sum)
}

func TestMulByMillisecond(t *testing.T) {
	sum, err := MulByMillisecond(math.MaxInt64 / time.Millisecond)
	require.NoError(t, err)
	require.Equal(t, time.Duration(9223372036854000000), sum)

	sum, err = MulByMillisecond(math.MaxInt64/time.Millisecond + 1)
	require.Error(t, err)
	require.Equal(t, time.Duration(0), sum)
}

func TestMulBySecond(t *testing.T) {
	sum, err := MulBySecond(math.MaxInt64 / time.Second)
	require.NoError(t, err)
	require.Equal(t, time.Duration(9223372036000000000), sum)

	sum, err = MulBySecond(math.MaxInt64/time.Second + 1)
	require.Error(t, err)
	require.Equal(t, time.Duration(0), sum)
}

func TestMulByMinute(t *testing.T) {
	sum, err := MulByMinute(math.MaxInt64 / time.Minute)
	require.NoError(t, err)
	require.Equal(t, time.Duration(9223372020000000000), sum)

	sum, err = MulByMinute(math.MaxInt64/time.Minute + 1)
	require.Error(t, err)
	require.Equal(t, time.Duration(0), sum)
}

func TestMulByHour(t *testing.T) {
	sum, err := MulByHour(math.MaxInt64 / time.Hour)
	require.NoError(t, err)
	require.Equal(t, time.Duration(9223369200000000000), sum)

	sum, err = MulByHour(math.MaxInt64/time.Hour + 1)
	require.Error(t, err)
	require.Equal(t, time.Duration(0), sum)
}
