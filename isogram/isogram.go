package isogram

import (
	"unicode"
)

func IsIsogram(word string) bool {
	var letter rune = 0
	for _, i := range word {
		if !unicode.IsLetter(i) {
			continue
		}
		j := rune(1 << (unicode.ToLower(i) - 'a'))
		if letter&j != 0 {
			return false
		}
		letter = letter | j
	}
	return true
}
