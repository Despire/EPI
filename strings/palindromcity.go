package strings

import (
	"unicode"
	"unicode/utf8"
)

func UTF8Palindrome(s string) bool {
	for utf8.RuneCountInString(s) > 1 {
		firstRune, firstSize := utf8.DecodeRuneInString(s)
		lastRune, lastSize := utf8.DecodeLastRuneInString(s)

		if unicode.ToLower(firstRune) != unicode.ToLower(lastRune) {
			return false
		}

		s = s[firstSize : len(s)-lastSize]
	}
	return true
}

func ASCIIPalindrome(s string) bool {
	i := 0
	j := len(s) - 1

	for i < j {
		for !(unicode.IsDigit(rune(s[i])) || unicode.IsLetter(rune(s[i]))) {
			i++
		}
		for !(unicode.IsDigit(rune(s[j])) || unicode.IsLetter(rune(s[j]))) {
			j--
		}

		letterI := unicode.ToLower(rune(s[i]))
		letterJ := unicode.ToLower(rune(s[j]))
		if letterI != letterJ {
			return false
		}

		i++
		j--
	}
	return true
}
