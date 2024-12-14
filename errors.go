package whilst

import "errors"

var (
	ErrNumberUnspecified = errors.New("number was not specified")
	ErrOnlyInteger       = errors.New("years, months and days can only be integer")
	ErrSymbolDotAgain    = errors.New("symbol of the dot was specified again")
	ErrSymbolSignAgain   = errors.New("symbol of the sign was specified again")
	ErrUnexpectedSymbol  = errors.New("unexpected symbol was specified")
	ErrUnexpectedUnit    = errors.New("unexpected unit was specified")
	ErrUnitUnspecified   = errors.New("unit was not specified")
)
