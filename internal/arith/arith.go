// Internal package used for integer calculations with overflow checking. It contains
// specialized functions for specific types, which work faster than generalized
// functions.
package arith

import (
	"time"

	"github.com/akramarenkov/safe"
	"github.com/akramarenkov/safe/intspec"
	"github.com/akramarenkov/whilst/internal/consts"
)

// Adds two integers of int64 type and detects whether an overflow has
// occurred or not.
//
// It is assumed that both numbers are always positive.
func Add(first int64, second int64) (int64, error) {
	sum := first + second

	if sum < first {
		return 0, safe.ErrOverflow
	}

	return sum, nil
}

// Adds two integers of time.Duration type and detects whether an overflow has
// occurred or not.
//
// It is assumed that both numbers are always positive.
func AddDuration(first time.Duration, second time.Duration) (time.Duration, error) {
	sum := first + second

	if sum < first {
		return 0, safe.ErrOverflow
	}

	return sum, nil
}

// Multiplies a number of int64 type by 10 and detects whether
// an overflow has occurred or not.
//
// It is assumed that the number is always positive.
func MulBy10(number int64) (int64, error) {
	const maximum = intspec.MaxInt64 / consts.DecimalBase

	if number > maximum {
		return 0, safe.ErrOverflow
	}

	return number * consts.DecimalBase, nil
}

// Multiplies a number of time.Duration type by time.Microsecond and detects whether
// an overflow has occurred or not.
//
// It is assumed that the number is always positive.
func MulByMicrosecond(number time.Duration) (time.Duration, error) {
	const maximum = intspec.MaxInt64 / time.Microsecond

	if number > maximum {
		return 0, safe.ErrOverflow
	}

	return number * time.Microsecond, nil
}

// Multiplies a number of time.Duration type by time.Millisecond and detects whether
// an overflow has occurred or not.
//
// It is assumed that the number is always positive.
func MulByMillisecond(number time.Duration) (time.Duration, error) {
	const maximum = intspec.MaxInt64 / time.Millisecond

	if number > maximum {
		return 0, safe.ErrOverflow
	}

	return number * time.Millisecond, nil
}

// Multiplies a number of time.Duration type by time.Second and detects whether
// an overflow has occurred or not.
//
// It is assumed that the number is always positive.
func MulBySecond(number time.Duration) (time.Duration, error) {
	const maximum = intspec.MaxInt64 / time.Second

	if number > maximum {
		return 0, safe.ErrOverflow
	}

	return number * time.Second, nil
}

// Multiplies a number of time.Duration type by time.Minute and detects whether
// an overflow has occurred or not.
//
// It is assumed that the number is always positive.
func MulByMinute(number time.Duration) (time.Duration, error) {
	const maximum = intspec.MaxInt64 / time.Minute

	if number > maximum {
		return 0, safe.ErrOverflow
	}

	return number * time.Minute, nil
}

// Multiplies a number of time.Duration type by time.Hour and detects whether
// an overflow has occurred or not.
//
// It is assumed that the number is always positive.
func MulByHour(number time.Duration) (time.Duration, error) {
	const maximum = intspec.MaxInt64 / time.Hour

	if number > maximum {
		return 0, safe.ErrOverflow
	}

	return number * time.Hour, nil
}
