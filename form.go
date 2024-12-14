package whilst

import (
	"time"

	"github.com/akramarenkov/whilst/internal/consts"
)

//nolint:gochecknoglobals // To increase the performance
var factors = [consts.FractionalLength]time.Duration{
	1e8, 1e7, 1e6, 1e5, 1e4, 1e3, 1e2, 1e1, 1e0,
}

func formFractional(output []byte, fractional time.Duration) []byte {
	if fractional == 0 {
		return output
	}

	output = append(output, consts.SymbolDot)

	fractional %= consts.FractionalFactor

	for id := range consts.FractionalLength {
		if fractional == 0 {
			break
		}

		digit := fractional / factors[id]
		fractional %= factors[id]

		output = append(output, '0'+byte(digit))
	}

	return output
}
