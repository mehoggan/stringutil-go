package stringutil

import "unicode"

// ToUpper uppercases all the runes in its argument string.
func ToLower(s string) string {
	r := []rune(s)
	for i := range r {
		r[i] = unicode.ToLower(r[i])
	}
	return string(r)
}
