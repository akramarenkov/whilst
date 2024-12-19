package whilst

import (
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	whl, err := Parse("0")
	require.NoError(t, err)
	require.Equal(t, "0s", whl.String())

	whl, err = Parse("-0")
	require.NoError(t, err)
	require.Equal(t, "0s", whl.String())

	whl, err = Parse("-0m")
	require.NoError(t, err)
	require.Equal(t, "0s", whl.String())

	whl, err = Parse("-2y3mo10d23.5h59.5m58s")
	require.NoError(t, err)
	require.Equal(t, "-2y3mo10d24h30m28s", whl.String())

	whl, err = Parse("-.5h.5m58s")
	require.NoError(t, err)
	require.Equal(t, "-31m28s", whl.String())

	whl, err = Parse("-2y  3mo10d23.5h59.5m58s")
	require.NoError(t, err)
	require.Equal(t, "-2y3mo10d24h30m28s", whl.String())

	whl, err = Parse("-2y3mo10d23.5h59.5m58s  ")
	require.NoError(t, err)
	require.Equal(t, "-2y3mo10d24h30m28s", whl.String())

	whl, err = Parse("10ms")
	require.NoError(t, err)
	require.Equal(t, "10ms", whl.String())

	whl, err = Parse("30us")
	require.NoError(t, err)
	require.Equal(t, "30µs", whl.String())

	whl, err = Parse("10ns")
	require.NoError(t, err)
	require.Equal(t, "10ns", whl.String())

	whl, err = Parse("65535d")
	require.NoError(t, err)
	require.Equal(t, "65535d", whl.String())

	whl, err = Parse("65535mo")
	require.NoError(t, err)
	require.Equal(t, "65535mo", whl.String())

	whl, err = Parse("65535y")
	require.NoError(t, err)
	require.Equal(t, "65535y", whl.String())

	whl, err = Parse("9223372036854775807ns")
	require.NoError(t, err)
	require.Equal(t, "2562047h47m16.854775807s", whl.String())

	whl, err = Parse("9223372036854775806ns1ns")
	require.NoError(t, err)
	require.Equal(t, "2562047h47m16.854775807s", whl.String())

	whl, err = Parse("-9223372036854775808ns")
	require.NoError(t, err)
	require.Equal(t, "-2562047h47m16.854775808s", whl.String())

	whl, err = Parse("-9223372036854775807ns1ns")
	require.NoError(t, err)
	require.Equal(t, "-2562047h47m16.854775808s", whl.String())
}

func TestParseError(t *testing.T) {
	inputs := []string{
		" - - 0s",
		" - + 0s",
		"-৩s",
		"-0.0.s",
		"-1",
		"-1 ",
		"1s s",
		"2.5y",
		"2.5mo",
		"2.5d",
		"2.5c",
		"18446744073709551616s",
		"184467440737095516100s",
		"18446744073709551615s0.1s",
		"18446744073709551615s.1s",
		"18446744073709551615s 1s",
		"18446744073709551615h",
		"18446744073709551615m",
		"18446744073709551615s",
		"18446744073709551615ms",
		"18446744073709551615µs",
		"9223372036854775807ns1ns",
		"-9223372036854775808ns1ns",
		"9223372036854775807ns0.1s",
		"-9223372036854775808ns0.1s",
		"65536d",
		"65535d1d",
		"65536mo",
		"65535mo1mo",
		"65536y",
		"65535y1y",
	}

	for _, input := range inputs {
		whl, err := Parse(input)
		require.Error(t, err, "input: %v", input)
		require.Equal(t, Whilst{}, whl, "input: %v", input)
	}
}

func TestCompatibility(t *testing.T) {
	inputs := []string{
		"-9223372036854775808ns",
		".10000000010000000000m",
		"0.20000000000000000001s",
		"0.9223372036854775808s",
	}

	for _, input := range inputs {
		whl, err := Parse(input)
		require.NoError(t, err, "input: %v", input)

		expected, err := time.ParseDuration(input)
		require.NoError(t, err, "input: %v", input)

		require.Equal(t, expected, whl.Duration(time.Time{}), "input: %v", input)
		require.Equal(t, expected.String(), whl.String(), "input: %v", input)
	}
}

func TestCompatibilityError(t *testing.T) {
	inputs := []string{
		"",
		"   ",
		"9223372036854775808ns",
		"-9223372036854775809ns",
	}

	for _, input := range inputs {
		_, err := Parse(input)
		require.Error(t, err, "input: %v", input)

		_, err = time.ParseDuration(input)
		require.Error(t, err, "input: %v", input)
	}
}

func TestManualSet(t *testing.T) {
	whl := Whilst{Nano: -1e9}
	require.Equal(t, "-1s", whl.String())
	require.Equal(t, "0000-12-31 23:59:59 +0000 UTC", whl.When(time.Time{}).String())

	whl = Whilst{Nano: 1e9, Negative: true}
	require.Equal(t, "-1s", whl.String())
	require.Equal(t, "0000-12-31 23:59:59 +0000 UTC", whl.When(time.Time{}).String())

	whl = Whilst{Nano: -1e9, Negative: true}
	require.Equal(t, "-1s", whl.String())
	require.Equal(t, "0000-12-31 23:59:59 +0000 UTC", whl.When(time.Time{}).String())

	whl = Whilst{Nano: 1e9}
	require.Equal(t, "1s", whl.String())
	require.Equal(t, "0001-01-01 00:00:01 +0000 UTC", whl.When(time.Time{}).String())
}

func FuzzPanic(f *testing.F) {
	f.Add(" - 2y 3mo 10d 23.5h 59.5m 58.01003001s 10ms 30µs 10ns")

	f.Fuzz(
		func(_ *testing.T, input string) {
			whl, err := Parse(input)
			if err != nil {
				return
			}

			_ = whl.String()
		},
	)
}

func FuzzDegradation(f *testing.F) {
	f.Add(" - 2y 3mo 10d 23.5h 59.5m 58.01003001s 10ms 30µs 10ns")

	f.Fuzz(
		func(t *testing.T, input string) {
			parsed1, err := Parse(input)
			if err != nil {
				return
			}

			formatted1 := parsed1.String()

			parsed2, err := Parse(formatted1)
			require.NoError(t, err)
			require.Equal(t, parsed1, parsed2)

			formatted2 := parsed2.String()
			require.Equal(t, formatted1, formatted2)
		},
	)
}

func FuzzManualSet(f *testing.F) {
	f.Add(
		int64(math.MaxInt64),
		uint16(math.MaxUint8),
		uint16(math.MaxUint8),
		uint16(math.MaxUint8),
		true,
	)

	f.Add(
		int64(math.MaxInt64),
		uint16(math.MaxUint8),
		uint16(math.MaxUint8),
		uint16(math.MaxUint8),
		false,
	)

	f.Add(
		int64(math.MinInt64),
		uint16(math.MaxUint8),
		uint16(math.MaxUint8),
		uint16(math.MaxUint8),
		true,
	)

	f.Add(
		int64(math.MinInt64),
		uint16(math.MaxUint8),
		uint16(math.MaxUint8),
		uint16(math.MaxUint8),
		false,
	)

	f.Fuzz(
		func(t *testing.T, nano int64, days, months, years uint16, negative bool) {
			origin := Whilst{
				Nano:     time.Duration(nano),
				Days:     days,
				Months:   months,
				Years:    years,
				Negative: negative,
			}

			formatted1 := origin.String()

			parsed, err := Parse(formatted1)
			require.NoError(t, err)
			require.Equal(t, origin.normalize(), parsed)

			formatted2 := parsed.String()
			require.Equal(t, formatted1, formatted2)
		},
	)
}

func FuzzCompatibility(f *testing.F) {
	f.Add("-23.5h59.5m58.01003001s10ms30µs10ns")
	f.Add("9223372036854775807ns")
	f.Add("-9223372036854775808ns")

	f.Fuzz(
		func(t *testing.T, input string) {
			expected, err := time.ParseDuration(input)
			if err != nil {
				return
			}

			whl, err := Parse(input)
			require.NoError(t, err)
			require.Equal(t, expected, whl.Duration(time.Time{}))
			require.Equal(t, expected.String(), whl.String())
		},
	)
}

func FuzzError(f *testing.F) {
	f.Add("-23.5h59.5m58.01003001s10ms30µs10ns")
	f.Add("9223372036854775807ns")
	f.Add("-9223372036854775808ns")

	f.Fuzz(
		func(t *testing.T, input string) {
			if _, err := Parse(input); err == nil {
				return
			}

			_, err := time.ParseDuration(input)
			require.Error(t, err)
		},
	)
}
