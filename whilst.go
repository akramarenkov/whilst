package whilst

import (
	"strconv"
	"time"

	"github.com/akramarenkov/whilst/internal/consts"
)

// Time duration with days, months and years.
//
// Maximum value of days, months and years is 65535.
type Whilst struct {
	duration uint64

	days   uint16
	months uint16
	years  uint16

	negative bool
}

// Parses a string representation of the duration.
//
// List of valid duration units:
//   - y            - year
//   - mo           - month
//   - d            - day
//   - h            - hour
//   - m            - minute
//   - s            - second
//   - ms           - millisecond
//   - µs | μs | us - microsecond
//   - ns           - nanosecond
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
		//nolint:gosec // Value of duration is controlled when parsing and setting
		// Value uint64(-MinInt64) when converted and inverted will
		// take the value int64(MinInt64)
		duration := -time.Duration(whl.duration)

		return from.AddDate(-int(whl.years), -int(whl.months), -int(whl.days)).Add(duration)
	}

	//nolint:gosec // Value of duration is controlled when parsing and setting
	duration := time.Duration(whl.duration)

	return from.AddDate(int(whl.years), int(whl.months), int(whl.days)).Add(duration)
}
