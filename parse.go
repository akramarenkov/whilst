package whilst

import (
	"time"

	"github.com/akramarenkov/safe"
	"github.com/akramarenkov/whilst/internal/ascii"
	"github.com/akramarenkov/whilst/internal/consts"
	"github.com/akramarenkov/whilst/internal/credible"
)

// Parsing context.
type parser struct {
	input string

	foundNum bool
	foundDot bool

	idUnit int

	integer  uint64
	fraction int64
	scale    float64

	whl *Whilst
}

// Parses the input string.
//
// It is assumed that the input string corresponds to the regular expression
// ^\s*[-+]?\s*([0-9]*(\.[0-9]*)?[a-z]+\s*)+$.
func parse(input string, whl *Whilst) error {
	prs := &parser{
		input: input,
		whl:   whl,
	}

	prs.reset()

	return prs.parse()
}

func (prs *parser) parse() error {
	if err := prs.begin(); err != nil {
		return err
	}

	if prs.input == specialZeroParse {
		prs.whl.negative = false
		return nil
	}

	for id, char := range []byte(prs.input) {
		if ascii.IsDigit(char) {
			if err := prs.onDigit(id); err != nil {
				return err
			}

			continue
		}

		if char == charDot {
			if err := prs.onDot(id); err != nil {
				return err
			}

			continue
		}

		if ascii.IsSpace(char) {
			if err := prs.onSpace(id); err != nil {
				return err
			}

			continue
		}

		if !prs.foundNum {
			return ErrNumberUnspecified
		}

		if prs.idUnit == -1 {
			prs.idUnit = id
		}
	}

	if err := prs.end(); err != nil {
		return err
	}

	if prs.whl.IsZero() {
		prs.whl.negative = false
	}

	return nil
}

func (prs *parser) begin() error {
	foundSign := false

	for id, char := range []byte(prs.input) {
		if char == charMinus || char == charPlus {
			if foundSign {
				return ErrCharSignAgain
			}

			foundSign = true

			prs.whl.negative = char == charMinus

			continue
		}

		if ascii.IsDigit(char) || char == charDot {
			prs.input = prs.input[id:]
			return nil
		}

		if ascii.IsSpace(char) {
			continue
		}

		return ErrUnexpectedChar
	}

	return ErrInputEmpty
}

func (prs *parser) onDigit(id int) error {
	char := prs.input[id]

	if !prs.foundNum {
		prs.foundNum = true
		return prs.incInteger(char)
	}

	if prs.idUnit != -1 {
		if err := prs.addValue(id); err != nil {
			return err
		}

		prs.foundNum = true
		prs.foundDot = false

		prs.reset()

		return prs.incInteger(char)
	}

	if prs.foundDot {
		prs.incFraction(char)
		return nil
	}

	return prs.incInteger(char)
}

func (prs *parser) onDot(id int) error {
	if !prs.foundNum {
		prs.foundNum = true
		prs.foundDot = true

		return nil
	}

	if prs.idUnit != -1 {
		if err := prs.addValue(id); err != nil {
			return err
		}

		prs.foundNum = true
		prs.foundDot = true

		prs.reset()

		return nil
	}

	if prs.foundDot {
		return ErrCharDotAgain
	}

	prs.foundDot = true

	return nil
}

func (prs *parser) onSpace(id int) error {
	if !prs.foundNum {
		return nil
	}

	if prs.idUnit == -1 {
		return ErrUnitUnspecified
	}

	if err := prs.addValue(id); err != nil {
		return err
	}

	prs.foundNum = false
	prs.foundDot = false

	prs.reset()

	return nil
}

func (prs *parser) end() error {
	if !prs.foundNum {
		return nil
	}

	if prs.idUnit == -1 {
		return ErrUnitUnspecified
	}

	return prs.addValue(len(prs.input))
}

func (prs *parser) incInteger(char byte) error {
	digit := ascii.ByteToDigit[uint64](char)

	shifted, err := credible.MulBy10U(prs.integer)
	if err != nil {
		return err
	}

	increased, err := safe.AddU(shifted, digit)
	if err != nil {
		return err
	}

	prs.integer = increased

	return nil
}

func (prs *parser) incFraction(char byte) {
	digit := ascii.ByteToDigit[int64](char)

	shifted, err := credible.MulBy10(prs.fraction)
	if err != nil {
		return
	}

	increased, err := credible.Add(shifted, digit)
	if err != nil {
		return
	}

	prs.fraction = increased
	prs.scale *= consts.DecimalBase
}

func (prs *parser) reset() {
	prs.idUnit = -1

	prs.integer = 0
	prs.fraction = 0
	prs.scale = 1
}

func (prs *parser) addValue(id int) error {
	whole := prs.integer
	dimension := time.Nanosecond
	unit := prs.input[prs.idUnit:id]

	switch unit {
	case unitYear:
		if prs.fraction != 0 {
			return ErrOnlyInteger
		}

		increased, err := credible.AddU64ToU16(prs.whl.years, prs.integer)
		if err != nil {
			return err
		}

		prs.whl.years = increased

		return nil
	case unitMonth:
		if prs.fraction != 0 {
			return ErrOnlyInteger
		}

		increased, err := credible.AddU64ToU16(prs.whl.months, prs.integer)
		if err != nil {
			return err
		}

		prs.whl.months = increased

		return nil
	case unitDay:
		if prs.fraction != 0 {
			return ErrOnlyInteger
		}

		increased, err := credible.AddU64ToU16(prs.whl.days, prs.integer)
		if err != nil {
			return err
		}

		prs.whl.days = increased

		return nil
	case unitHour:
		dimension = time.Hour

		multiplied, err := credible.MulByHour(whole)
		if err != nil {
			return err
		}

		whole = multiplied
	case unitMinute:
		dimension = time.Minute

		multiplied, err := credible.MulByMinute(whole)
		if err != nil {
			return err
		}

		whole = multiplied
	case unitSecond:
		dimension = time.Second

		multiplied, err := credible.MulBySecond(whole)
		if err != nil {
			return err
		}

		whole = multiplied
	case unitMillisecond:
		dimension = time.Millisecond

		multiplied, err := credible.MulByMillisecond(whole)
		if err != nil {
			return err
		}

		whole = multiplied
	case unitMicrosecond, unitMicrosecondA1, unitMicrosecondA2:
		dimension = time.Microsecond

		multiplied, err := credible.MulByMicrosecond(whole)
		if err != nil {
			return err
		}

		whole = multiplied
	case unitNanosecond:
	default:
		return ErrUnexpectedUnit
	}

	duration, err := credible.AddDuration(prs.whl.duration, whole, prs.whl.negative)
	if err != nil {
		return err
	}

	if prs.fraction == 0 {
		prs.whl.duration = duration
		return nil
	}

	converted := float64(prs.fraction) * (float64(dimension) / float64(prs.scale))

	duration, err = credible.AddDuration(duration, uint64(converted), prs.whl.negative)
	if err != nil {
		return err
	}

	prs.whl.duration = duration

	return nil
}
