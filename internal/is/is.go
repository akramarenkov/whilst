// Internal package used to check whether symbols belong to certain categories.
package is

const (
	symbolNBSP = 0xA0
	symbolNEL  = 0x85
)

// Reports whether the symbol belongs to whitespace characters.
func Space(symbol byte) bool {
	switch symbol {
	case '\t', '\n', '\v', '\f', '\r', ' ', symbolNBSP, symbolNEL:
		return true
	}

	return false
}

// Reports whether the symbol belongs to decimal digital characters.
func Digit(symbol byte) bool {
	return symbol >= '0' && symbol <= '9'
}
