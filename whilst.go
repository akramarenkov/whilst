package whilst

import (
	"strconv"
	"time"

	"github.com/akramarenkov/whilst/internal/consts"
)

// Time duration with days, months and years.
type Whilst struct {
	duration uint64

	days   uint16
	months uint16
	years  uint16

	negative bool
}

// Parses a string representation of the duration.
//
// A duration string consists of several numbers supplemented with a unit of
// measurement. There may be spaces between a numbers supplemented with a unit,
// but there must not be spaces between a number and a unit. One of a signs - or + can
// be specified at a beginning of a string.
//
// A value of days, months and years can only be an integer and cannot be greater
// than 65535 for each.
//
// Remaining values ​​must not be greater than 9223372036854775807 for positive duration
// and 9223372036854775808 for negative duration and may have a fractional part.
//
// List of valid units:
//   - y            - year
//   - mo           - month
//   - d            - day
//   - h            - hour
//   - m            - minute
//   - s            - second
//   - ms           - millisecond
//   - µs | μs | us - microsecond
//   - ns           - nanosecond
//
// Example of strings:
//   - 2y3mo10d 24h30m28.02006002s
//   - - 2y3mo10d24h30m28.02006002s
//   - + 2y 3mo 10d 24h 30m 28.02006002s
func Parse(input string) (Whilst, error) {
	whl := Whilst{}

	if err := parse(input, &whl); err != nil {
		return Whilst{}, err
	}

	return whl, nil
}

// Reports whether the duration is zero.
func (whl Whilst) IsZero() bool {
	return whl.years|whl.months|whl.days == 0 && whl.duration == 0
}

// Returns a string representation of the duration.
func (whl Whilst) String() string {
	if whl.IsZero() {
		return specialZeroFormat
	}

	var output []byte

	switch {
	case whl.years|whl.months|whl.days == 0:
		output = make([]byte, 0, len(formatMaximumStd))
	default:
		output = make([]byte, 0, len(formatMaximum))
	}

	if whl.negative {
		output = append(output, charMinus)
	}

	if whl.years != 0 {
		output = strconv.AppendUint(output, uint64(whl.years), consts.DecimalBase)
		output = append(output, unitYear...)
	}

	if whl.months != 0 {
		output = strconv.AppendUint(output, uint64(whl.months), consts.DecimalBase)
		output = append(output, unitMonth...)
	}

	if whl.days != 0 {
		output = strconv.AppendUint(output, uint64(whl.days), consts.DecimalBase)
		output = append(output, unitDay...)
	}

	output = whl.appendHMS(output)

	return string(output)
}

func (whl Whilst) appendHMS(output []byte) []byte {
	duration := whl.duration
	upper := false

	hours := duration / consts.U64Hour
	duration %= consts.U64Hour

	minutes := duration / consts.U64Minute
	duration %= consts.U64Minute

	seconds := duration / consts.U64Second
	duration %= consts.U64Second

	if hours != 0 {
		upper = true

		output = strconv.AppendUint(output, hours, consts.DecimalBase)
		output = append(output, unitHour...)
	}

	if minutes != 0 || upper {
		upper = true

		output = strconv.AppendUint(output, minutes, consts.DecimalBase)
		output = append(output, unitMinute...)
	}

	if seconds != 0 || upper {
		output = strconv.AppendUint(output, seconds, consts.DecimalBase)
		output = appendFraction(output, duration)
		output = append(output, unitSecond...)

		return output
	}

	milliseconds := duration / consts.U64Millisecond
	duration %= consts.U64Millisecond
	millisecondsFraction := duration * consts.U64Second / consts.U64Millisecond

	if milliseconds != 0 {
		output = strconv.AppendUint(output, milliseconds, consts.DecimalBase)
		output = appendFraction(output, millisecondsFraction)
		output = append(output, unitMillisecond...)

		return output
	}

	microseconds := duration / consts.U64Microsecond
	duration %= consts.U64Microsecond
	microsecondsFraction := duration * consts.U64Second / consts.U64Microsecond

	if microseconds != 0 {
		output = strconv.AppendUint(output, microseconds, consts.DecimalBase)
		output = appendFraction(output, microsecondsFraction)
		output = append(output, unitMicrosecond...)

		return output
	}

	if duration != 0 {
		output = strconv.AppendUint(output, duration, consts.DecimalBase)
		output = append(output, unitNanosecond...)
	}

	return output
}

// Returns a time.Duration representation of the duration.
//
// Time from is necessary because shift by days, months and years is not deterministic
// and depends on the time relative to which it occurs.
func (whl Whilst) Duration(from time.Time) time.Duration {
	return whl.When(from).Sub(from)
}

// Returns a time shifted by the duration.
func (whl Whilst) When(from time.Time) time.Time {
	if whl.negative {
		// Value uint64(-MinInt64) when converted to int64 will take the value MinInt64,
		// the value int64(MinInt64) when inverted will also take the value MinInt64
		//
		//nolint:gosec // Value of duration is controlled when parsing and setting
		duration := -time.Duration(whl.duration)

		return from.AddDate(-int(whl.years), -int(whl.months), -int(whl.days)).Add(duration)
	}

	//nolint:gosec // Value of duration is controlled when parsing and setting
	duration := time.Duration(whl.duration)

	return from.AddDate(int(whl.years), int(whl.months), int(whl.days)).Add(duration)
}
