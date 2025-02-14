package whilst

import "github.com/akramarenkov/whilst/internal/ascii"

const (
	fractionLength = 9
)

//nolint:gochecknoglobals // To increase performance
var dividers = [fractionLength]uint64{
	1e8, 1e7, 1e6, 1e5, 1e4, 1e3, 1e2, 1e1, 1e0,
}

// Length of a fraction value in decimal notation must not exceed fractionLength.
func appendFraction(output []byte, fraction uint64) []byte {
	if fraction == 0 {
		return output
	}

	output = append(output, charDot)

	for _, divider := range dividers {
		digit := fraction / divider
		fraction %= divider

		output = append(output, ascii.DigitToByte(digit))

		if fraction == 0 {
			break // For coverage
		}
	}

	return output
}
