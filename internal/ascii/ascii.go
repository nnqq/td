// Package ascii provides data and functions to test some properties of
// ASCII code points.
package ascii

// IsDigit reports whether the rune is a decimal digit.
func IsDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

// IsLatinLetter reports whether the rune is ASCII latin letter.
func IsLatinLetter(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}
