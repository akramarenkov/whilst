package whilst

import (
	"time"

	"github.com/akramarenkov/whilst/internal/arith"
	"github.com/akramarenkov/whilst/internal/consts"
	"github.com/akramarenkov/whilst/internal/is"
)

type parser struct {
	input string

	foundNum bool
	foundDot bool

	idUnit int

	integer    int64
	fractional int64
	scale      float64

	whl Whilst
}

func newParser(input string) parser {
	prs := parser{
		input:  input,
		idUnit: -1,
		scale:  1,
	}

	return prs
}

// Corresponds to the regular expression ^\s*[-+]?\s*([0-9]*(\.[0-9]*)?[a-z]+\s*)+$.
func (prs *parser) parse() (Whilst, error) {
	if err := prs.begin(); err != nil {
		return Whilst{}, err
	}

	if prs.input == "0" {
		if prs.whl.isZero() {
			prs.whl.negative = false
		}

		return prs.whl, nil
	}

	for id, symbol := range []byte(prs.input) {
		if is.Digit(symbol) {
			if err := prs.onDigit(id); err != nil {
				return Whilst{}, err
			}

			continue
		}

		if symbol == consts.SymbolDot {
			if err := prs.onDot(id); err != nil {
				return Whilst{}, err
			}

			continue
		}

		if is.Space(symbol) {
			if err := prs.onSpace(id); err != nil {
				return Whilst{}, err
			}

			continue
		}

		if !prs.foundNum {
			return Whilst{}, ErrNumberUnspecified
		}

		if prs.idUnit == -1 {
			prs.idUnit = id
		}
	}

	if err := prs.end(); err != nil {
		return Whilst{}, err
	}

	if prs.whl.isZero() {
		prs.whl.negative = false
	}

	return prs.whl, nil
}

func (prs *parser) begin() error {
	foundSign := false

	for id, symbol := range []byte(prs.input) {
		if symbol == consts.SymbolMinus || symbol == consts.SymbolPlus {
			if foundSign {
				return ErrSymbolSignAgain
			}

			foundSign = true

			prs.whl.negative = symbol == consts.SymbolMinus

			continue
		}

		if is.Digit(symbol) || symbol == consts.SymbolDot {
			prs.input = prs.input[id:]
			return nil
		}

		if is.Space(symbol) {
			continue
		}

		return ErrUnexpectedSymbol
	}

	return nil
}

func (prs *parser) onDigit(id int) error {
	symbol := prs.input[id]

	if !prs.foundNum {
		prs.foundNum = true
		return prs.increaseInteger(symbol)
	}

	if prs.idUnit != -1 {
		if err := prs.addValue(id); err != nil {
			return err
		}

		prs.foundNum = true
		prs.foundDot = false

		prs.reset()

		return prs.increaseInteger(symbol)
	}

	if prs.foundDot {
		prs.increaseFractional(symbol)
		return nil
	}

	return prs.increaseInteger(symbol)
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
		return ErrSymbolDotAgain
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

func (prs *parser) increaseInteger(symbol byte) error {
	digit := int64(symbol - '0')

	shifted, err := arith.MulBy10(prs.integer)
	if err != nil {
		return err
	}

	increased, err := arith.Add(shifted, digit)
	if err != nil {
		return err
	}

	prs.integer = increased

	return nil
}

func (prs *parser) increaseFractional(symbol byte) {
	digit := int64(symbol - '0')

	shifted, err := arith.MulBy10(prs.fractional)
	if err != nil {
		return
	}

	increased, err := arith.Add(shifted, digit)
	if err != nil {
		return
	}

	prs.fractional = increased
	prs.scale *= consts.DecimalBase
}

func (prs *parser) reset() {
	prs.idUnit = -1

	prs.integer = 0
	prs.fractional = 0
	prs.scale = 1
}

func (prs *parser) addValue(id int) error {
	whole := time.Duration(prs.integer)
	dimension := time.Nanosecond
	unit := prs.input[prs.idUnit:id]

	switch unit {
	case "y":
		if prs.fractional != 0 {
			return ErrOnlyInteger
		}

		increased, err := arith.Add(prs.whl.years, prs.integer)
		if err != nil {
			return err
		}

		prs.whl.years = increased

		return nil
	case "mo":
		if prs.fractional != 0 {
			return ErrOnlyInteger
		}

		increased, err := arith.Add(prs.whl.months, prs.integer)
		if err != nil {
			return err
		}

		prs.whl.months = increased

		return nil
	case "d":
		if prs.fractional != 0 {
			return ErrOnlyInteger
		}

		increased, err := arith.Add(prs.whl.days, prs.integer)
		if err != nil {
			return err
		}

		prs.whl.days = increased

		return nil
	case "h":
		dimension = time.Hour

		multiplied, err := arith.MulByHour(whole)
		if err != nil {
			return err
		}

		whole = multiplied
	case "m":
		dimension = time.Minute

		multiplied, err := arith.MulByMinute(whole)
		if err != nil {
			return err
		}

		whole = multiplied
	case "s":
		dimension = time.Second

		multiplied, err := arith.MulBySecond(whole)
		if err != nil {
			return err
		}

		whole = multiplied
	case "ms":
		dimension = time.Millisecond

		multiplied, err := arith.MulByMillisecond(whole)
		if err != nil {
			return err
		}

		whole = multiplied
	case "µs", "μs", "us":
		dimension = time.Microsecond

		multiplied, err := arith.MulByMicrosecond(whole)
		if err != nil {
			return err
		}

		whole = multiplied
	case "ns":
	default:
		return ErrUnexpectedUnit
	}

	duration, err := arith.AddDuration(prs.whl.duration, whole)
	if err != nil {
		return err
	}

	if prs.fractional == 0 {
		prs.whl.duration = duration
		return nil
	}

	converted := float64(prs.fractional) * (float64(dimension) / float64(prs.scale))

	duration, err = arith.AddDuration(duration, time.Duration(converted))
	if err != nil {
		return err
	}

	prs.whl.duration = duration

	return nil
}
