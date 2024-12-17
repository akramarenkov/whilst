// Internal package used to work with 7-bit US-ASCII characters.
package ascii

import "golang.org/x/exp/constraints"

// Reports whether the byte value belongs to whitespace characters.
func IsSpace(value byte) bool {
	switch value {
	case '\t', '\n', '\v', '\f', '\r', ' ':
		return true
	}

	return false
}

// Reports whether the byte value belongs to decimal digital characters.
func IsDigit(value byte) bool {
	return value >= '0' && value <= '9'
}

// Converts the byte value belongs to decimal digital characters to a digit of the
// specified integer type.
//
// Byte value correctness checks are not performed.
func ByteToDigit[Type constraints.Integer](value byte) Type {
	return Type(value - '0')
}

// Converts the digit value to decimal digital character.
//
// Digit value correctness checks are not performed.
func DigitToByte[Type constraints.Integer](digit Type) byte {
	return '0' + byte(digit)
}
