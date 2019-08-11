// Copyright (c) 2019 Matúš Mrekaj.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package strings

import (
	"unicode"
	"unicode/utf8"
)

// UTF8Palindrome takes as input a utf8 enconded string
// and checks if it is a palindrome.
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

// ASCIIPalindrome takes as input a string
// and checks if it is a palindrome.
// Takes O(N) time with O(1) space complexity.
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
