package cipher

import (
	"math"
	"strings"
	"unicode"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// Performs the atbash cipher on a given input string and returns the encoded string.
func Atbash(s string) string {
	var encoded string

	for _, token := range s {
		// find the alphabetical index of current input token
		index := strings.Index(alphabet, string(unicode.ToLower(token)))

		// not an alphabetical token
		if index == -1 {
			encoded += string(token)
			continue
		}

		// get the symmetrical opposite position in the alphabet for the given token
		index = int(math.Abs(float64(index - 25)))

		if unicode.IsLower(token) {
			encoded += string(alphabet[index])
		} else {
			encoded += string(unicode.ToUpper(rune(alphabet[index])))
		}
	}

	return encoded
}
