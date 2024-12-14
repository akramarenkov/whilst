package whilst

import (
	"strconv"
	"time"

	"github.com/akramarenkov/whilst/internal/consts"
)

// Time duration with days, months and years.
type Whilst struct {
	negative bool

	duration time.Duration

	days   int64
	months int64
	years  int64
}

// Parses a string representation of a duration.
func Parse(input string) (Whilst, error) {
	parser := newParser(input)
	return parser.parse()
}

// Returns a string representation of a duration.
func (whl Whilst) String() string {
	if whl.isZero() {
		return "0s"
	}

	output := make([]byte, 0, len("-2562047h47m16.854775808s"))

	if whl.negative {
		output = append(output, consts.SymbolMinus)
	}

	if whl.years != 0 {
		output = strconv.AppendInt(output, whl.years, consts.DecimalBase)
		output = append(output, "y"...)
	}

	if whl.months != 0 {
		output = strconv.AppendInt(output, whl.months, consts.DecimalBase)
		output = append(output, "mo"...)
	}

	if whl.days != 0 {
		output = strconv.AppendInt(output, whl.days, consts.DecimalBase)
		output = append(output, "d"...)
	}

	output = whl.formHMS(output)

	return string(output)
}

func (whl Whilst) isZero() bool {
	return whl.years|whl.months|whl.days|int64(whl.duration) == 0
}

func (whl Whilst) formHMS(output []byte) []byte {
	duration := whl.duration
	upper := false

	hours := duration / time.Hour
	duration %= time.Hour

	minutes := duration / time.Minute
	duration %= time.Minute

	seconds := duration / time.Second
	duration %= time.Second

	if hours != 0 || upper {
		upper = true

		output = strconv.AppendInt(output, int64(hours), consts.DecimalBase)
		output = append(output, "h"...)
	}

	if minutes != 0 || upper {
		upper = true

		output = strconv.AppendInt(output, int64(minutes), consts.DecimalBase)
		output = append(output, "m"...)
	}

	if seconds != 0 || upper {
		output = strconv.AppendInt(output, int64(seconds), consts.DecimalBase)
		output = formFractional(output, duration)
		output = append(output, "s"...)

		return output
	}

	milliseconds := duration / time.Millisecond
	duration %= time.Millisecond
	millisecondsFractional := duration * time.Second / time.Millisecond

	microseconds := duration / time.Microsecond
	duration %= time.Microsecond
	microsecondsFractional := duration * time.Second / time.Microsecond

	if milliseconds != 0 {
		output = strconv.AppendInt(output, int64(milliseconds), consts.DecimalBase)
		output = formFractional(output, millisecondsFractional)
		output = append(output, "ms"...)

		return output
	}

	if microseconds != 0 {
		output = strconv.AppendInt(output, int64(microseconds), consts.DecimalBase)
		output = formFractional(output, microsecondsFractional)
		output = append(output, "µs"...)

		return output
	}

	if duration != 0 {
		output = strconv.AppendInt(output, int64(duration), consts.DecimalBase)
		output = append(output, "ns"...)
	}

	return output
}

// Returns a time.Duration representation of a duration.
//
// Time from is necessary because shift by days, months and years is not deterministic
// and depends on the time relative to which it occurs.
func (whl Whilst) Duration(from time.Time) time.Duration {
	return whl.When(from).Sub(from)
}

// Returns a time shifted by a duration.
func (whl Whilst) When(from time.Time) time.Time {
	if whl.negative {
		return from.AddDate(int(-whl.years), int(-whl.months), int(-whl.days)).Add(-whl.duration)
	}

	return from.AddDate(int(whl.years), int(whl.months), int(whl.days)).Add(whl.duration)
}
