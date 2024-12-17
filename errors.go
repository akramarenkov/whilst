package whilst

import "errors"

var (
	ErrCharDotAgain      = errors.New("character of the dot was specified again")
	ErrCharSignAgain     = errors.New("character of the sign was specified again")
	ErrInputEmpty        = errors.New("input string is empty")
	ErrNumberUnspecified = errors.New("number was not specified")
	ErrOnlyInteger       = errors.New("years, months and days can only be integer")
	ErrUnexpectedChar    = errors.New("unexpected character was specified")
	ErrUnexpectedUnit    = errors.New("unexpected unit was specified")
	ErrUnitUnspecified   = errors.New("unit was not specified")
)
