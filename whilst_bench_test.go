package whilst

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const (
	benchmarkInput            = " - 2y 3mo 10d 23.5h 59.5m 58.01003001s 10ms 30µs 10ns"
	benchmarkExpected         = "-2y3mo10d24h30m28.02006002s"
	benchmarkDurationInput    = "-23.5h59.5m58.01003001s10ms30µs10ns"
	benchmarkDurationExpected = "-24h30m28.02006002s"
)

func BenchmarkParseReference(b *testing.B) {
	var last int

	for range b.N {
		for id := range []byte(benchmarkInput) {
			last = id
		}
	}

	require.NotZero(b, last)
}

func BenchmarkParse(b *testing.B) {
	var (
		whl Whilst
		err error
	)

	for range b.N {
		whl, err = Parse(benchmarkInput)
	}

	require.NoError(b, err)
	require.Equal(b, benchmarkExpected, whl.String())
}

func BenchmarkString(b *testing.B) {
	whl, err := Parse(benchmarkInput)
	require.NoError(b, err)

	var output string

	for range b.N {
		output = whl.String()
	}

	require.Equal(b, benchmarkExpected, output)
}

func BenchmarkParseDurationReference(b *testing.B) {
	var last int

	for range b.N {
		for id := range []byte(benchmarkDurationInput) {
			last = id
		}
	}

	require.NotZero(b, last)
}

func BenchmarkParseDuration(b *testing.B) {
	var (
		whl Whilst
		err error
	)

	for range b.N {
		whl, err = Parse(benchmarkDurationInput)
	}

	require.NoError(b, err)
	require.Equal(b, benchmarkDurationExpected, whl.String())
}

func BenchmarkParseDurationStd(b *testing.B) {
	var (
		duration time.Duration
		err      error
	)

	for range b.N {
		duration, err = time.ParseDuration(benchmarkDurationInput)
	}

	require.NoError(b, err)
	require.Equal(b, benchmarkDurationExpected, duration.String())
}

func BenchmarkStringDuration(b *testing.B) {
	whl, err := Parse(benchmarkDurationInput)
	require.NoError(b, err)

	var output string

	for range b.N {
		output = whl.String()
	}

	require.Equal(b, benchmarkDurationExpected, output)
}

func BenchmarkStringDurationStd(b *testing.B) {
	duration, err := time.ParseDuration(benchmarkDurationInput)
	require.NoError(b, err)

	var output string

	for range b.N {
		output = duration.String()
	}

	require.Equal(b, benchmarkDurationExpected, output)
}
