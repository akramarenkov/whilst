package whilst

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func FuzzPanic(f *testing.F) {
	f.Add("-2y3mo10d23h59m58.01003001s")

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
	f.Add("-2y3mo10d23h59m58.01003001s")

	f.Fuzz(
		func(t *testing.T, input string) {
			parsed1, err := Parse(input)
			if err != nil {
				return
			}

			string1 := parsed1.String()

			parsed2, err := Parse(string1)
			require.NoError(t, err)
			require.Equal(t, parsed1, parsed2)

			string2 := parsed2.String()
			require.Equal(t, string1, string2)
		},
	)
}

func FuzzCompatibility(f *testing.F) {
	f.Add("-23h59m58.01003001s")

	f.Fuzz(
		func(t *testing.T, input string) {
			duration, err := time.ParseDuration(input)
			if err != nil {
				return
			}

			whl, err := Parse(input)
			require.NoError(t, err)
			require.Equal(t, duration, whl.Duration(time.Time{}))
			require.Equal(t, duration.String(), whl.String())
		},
	)
}

func FuzzError(f *testing.F) {
	f.Add("-23h59m58.01003001s")

	f.Fuzz(
		func(t *testing.T, input string) {
			_, err := Parse(input)
			if err == nil {
				return
			}

			_, err = time.ParseDuration(input)
			require.Error(t, err)
		},
	)
}
