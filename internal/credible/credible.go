// Internal package used for integer calculations with overflow checking. It contains
// specialized functions for specific types, which work faster than generalized
// functions.
package credible

import (
	"github.com/akramarenkov/safe"
	"github.com/akramarenkov/safe/intspec"
	"github.com/akramarenkov/whilst/internal/consts"
)

// Adds two integers of int64 type and detects whether
// an overflow has occurred or not.
//
// It is assumed that both integers are always positive.
func Add(first int64, second int64) (int64, error) {
	sum := first + second

	if sum < first {
		return 0, safe.ErrOverflow
	}

	return sum, nil
}

// Adds integer of uint64 type to integer of uint16 type and detects whether
// an overflow has occurred or not.
func AddU64ToU16(first uint16, second uint64) (uint16, error) {
	if second > intspec.MaxUint16 {
		return 0, safe.ErrOverflow
	}

	sum := first + uint16(second)

	if sum < first {
		return 0, safe.ErrOverflow
	}

	return sum, nil
}

// Adds two integers of uint64 type, detects whether an overflow has occurred or not and
// also detects whether the sum exceeds the maximum value for the time.Duration type in
// the uint64 representation.
func AddDuration(first uint64, second uint64, negative bool) (uint64, error) {
	const (
		minimum = -intspec.MinInt64
		maximum = intspec.MaxInt64
	)

	sum := first + second

	if sum < first {
		return 0, safe.ErrOverflow
	}

	switch negative {
	case false:
		if sum > maximum {
			return 0, safe.ErrOverflow
		}
	case true:
		if sum > minimum {
			return 0, safe.ErrOverflow
		}
	}

	return sum, nil
}

// Multiplies a integer of int64 type by 10 and detects whether
// an overflow has occurred or not.
//
// It is assumed that the integer is always positive.
func MulBy10(number int64) (int64, error) {
	const maximum = intspec.MaxInt64 / consts.DecimalBase

	if number > maximum {
		return 0, safe.ErrOverflow
	}

	return number * consts.DecimalBase, nil
}

// Multiplies a integer of uint64 type by 10 and detects whether
// an overflow has occurred or not.
func MulBy10U(number uint64) (uint64, error) {
	const maximum = intspec.MaxUint64 / consts.DecimalBase

	if number > maximum {
		return 0, safe.ErrOverflow
	}

	return number * consts.DecimalBase, nil
}

// Multiplies a integer of uint64 type by time.Microsecond and detects whether
// an overflow has occurred or not.
func MulByMicrosecond(number uint64) (uint64, error) {
	const maximum = intspec.MaxUint64 / consts.U64Microsecond

	if number > maximum {
		return 0, safe.ErrOverflow
	}

	return number * consts.U64Microsecond, nil
}

// Multiplies a integer of uint64 type by time.Millisecond and detects whether
// an overflow has occurred or not.
func MulByMillisecond(number uint64) (uint64, error) {
	const maximum = intspec.MaxUint64 / consts.U64Millisecond

	if number > maximum {
		return 0, safe.ErrOverflow
	}

	return number * consts.U64Millisecond, nil
}

// Multiplies a integer of uint64 type by time.Second and detects whether
// an overflow has occurred or not.
func MulBySecond(number uint64) (uint64, error) {
	const maximum = intspec.MaxUint64 / consts.U64Second

	if number > maximum {
		return 0, safe.ErrOverflow
	}

	return number * consts.U64Second, nil
}

// Multiplies a integer of uint64 type by time.Minute and detects whether
// an overflow has occurred or not.
func MulByMinute(number uint64) (uint64, error) {
	const maximum = intspec.MaxUint64 / consts.U64Minute

	if number > maximum {
		return 0, safe.ErrOverflow
	}

	return number * consts.U64Minute, nil
}

// Multiplies a integer of uint64 type by time.Hour and detects whether
// an overflow has occurred or not.
func MulByHour(number uint64) (uint64, error) {
	const maximum = intspec.MaxUint64 / consts.U64Hour

	if number > maximum {
		return 0, safe.ErrOverflow
	}

	return number * consts.U64Hour, nil
}
