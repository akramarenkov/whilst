package credible

import (
	"math"
	"testing"

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

func TestAddU64ToU16(t *testing.T) {
	sum, err := AddU64ToU16(math.MaxUint16, 0)
	require.NoError(t, err)
	require.Equal(t, uint16(math.MaxUint16), sum)

	sum, err = AddU64ToU16(0, math.MaxUint16)
	require.NoError(t, err)
	require.Equal(t, uint16(math.MaxUint16), sum)

	sum, err = AddU64ToU16(math.MaxUint16, 1)
	require.Error(t, err)
	require.Equal(t, uint16(0), sum)

	sum, err = AddU64ToU16(0, math.MaxUint16+1)
	require.Error(t, err)
	require.Equal(t, uint16(0), sum)
}

func TestAddU64ToS64(t *testing.T) {
	sum, err := AddU64ToS64(math.MaxInt64, 0, false)
	require.NoError(t, err)
	require.Equal(t, int64(math.MaxInt64), sum)

	sum, err = AddU64ToS64(0, math.MaxInt64, false)
	require.NoError(t, err)
	require.Equal(t, int64(math.MaxInt64), sum)

	sum, err = AddU64ToS64(0, -math.MinInt64, true)
	require.NoError(t, err)
	require.Equal(t, int64(math.MinInt64), sum)

	sum, err = AddU64ToS64(math.MaxInt64, 1, false)
	require.Error(t, err)
	require.Equal(t, int64(0), sum)

	sum, err = AddU64ToS64(0, math.MaxInt64+1, false)
	require.Error(t, err)
	require.Equal(t, int64(0), sum)

	sum, err = AddU64ToS64(math.MinInt64, 1, true)
	require.Error(t, err)
	require.Equal(t, int64(0), sum)

	sum, err = AddU64ToS64(0, math.MaxInt64+2, true)
	require.Error(t, err)
	require.Equal(t, int64(0), sum)

	sum, err = AddU64ToS64(-1, math.MaxInt64+1, true)
	require.Error(t, err)
	require.Equal(t, int64(0), sum)
}

func TestMulBy10(t *testing.T) {
	sum, err := MulBy10(math.MaxInt64 / consts.DecimalBase)
	require.NoError(t, err)
	require.Equal(t, int64(9223372036854775800), sum)

	sum, err = MulBy10(math.MaxInt64/consts.DecimalBase + 1)
	require.Error(t, err)
	require.Equal(t, int64(0), sum)
}

func TestMulBy10U(t *testing.T) {
	sum, err := MulBy10U(math.MaxUint64 / consts.DecimalBase)
	require.NoError(t, err)
	require.Equal(t, uint64(18446744073709551610), sum)

	sum, err = MulBy10U(math.MaxUint64/consts.DecimalBase + 1)
	require.Error(t, err)
	require.Equal(t, uint64(0), sum)
}

func TestMulByMicrosecond(t *testing.T) {
	sum, err := MulByMicrosecond(math.MaxUint64 / consts.U64Microsecond)
	require.NoError(t, err)
	require.Equal(t, uint64(18446744073709551000), sum)

	sum, err = MulByMicrosecond(math.MaxUint64/consts.U64Microsecond + 1)
	require.Error(t, err)
	require.Equal(t, uint64(0), sum)
}

func TestMulByMillisecond(t *testing.T) {
	sum, err := MulByMillisecond(math.MaxUint64 / consts.U64Millisecond)
	require.NoError(t, err)
	require.Equal(t, uint64(18446744073709000000), sum)

	sum, err = MulByMillisecond(math.MaxUint64/consts.U64Millisecond + 1)
	require.Error(t, err)
	require.Equal(t, uint64(0), sum)
}

func TestMulBySecond(t *testing.T) {
	sum, err := MulBySecond(math.MaxUint64 / consts.U64Second)
	require.NoError(t, err)
	require.Equal(t, uint64(18446744073000000000), sum)

	sum, err = MulBySecond(math.MaxUint64/consts.U64Second + 1)
	require.Error(t, err)
	require.Equal(t, uint64(0), sum)
}

func TestMulByMinute(t *testing.T) {
	sum, err := MulByMinute(math.MaxUint64 / consts.U64Minute)
	require.NoError(t, err)
	require.Equal(t, uint64(18446744040000000000), sum)

	sum, err = MulByMinute(math.MaxUint64/consts.U64Minute + 1)
	require.Error(t, err)
	require.Equal(t, uint64(0), sum)
}

func TestMulByHour(t *testing.T) {
	sum, err := MulByHour(math.MaxUint64 / consts.U64Hour)
	require.NoError(t, err)
	require.Equal(t, uint64(18446742000000000000), sum)

	sum, err = MulByHour(math.MaxUint64/consts.U64Hour + 1)
	require.Error(t, err)
	require.Equal(t, uint64(0), sum)
}
