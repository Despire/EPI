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
	"bytes"
	"unicode"
)

// charsmap maps hex characters to numeric values.
var charsmap = map[rune]int{
	'A': 10,
	'B': 11,
	'C': 12,
	'D': 13,
	'E': 14,
	'F': 15,
}

// decimap maps numeric values to hex characters.
var decimap = map[int]rune{
	10: 'A',
	11: 'B',
	12: 'C',
	13: 'D',
	14: 'E',
	15: 'F',
}

// convertBase converts the number in base 1 to another
// number from base 2.
func convertBase(number string, b1 int, b2 int) string {
	prefix := ""

	switch number[0] {
	case '+':
		prefix = "+"
	case '-':
		prefix = "-"
	}

	// strip the prefix
	if prefix != "" {
		number = number[1:]
	}

	base10 := toBase10(number, b1)

	return prefix + convertFromBase10(base10, b2)
}

// toBase10 converts a non-negative number from b1 to a number from base10.
// Takes O(N) time with O(1) space complexity.
func toBase10(number string, b1 int) int {
	base10 := 0

	for _, digit := range number {
		base10 *= b1

		if unicode.IsDigit(digit) {
			base10 += int(digit - '0')
		} else {
			base10 += charsmap[digit]
		}
	}

	return base10
}

// convertFromBase10 converts a non-negative number from base 10
// to a number from the specified base in the parameter b2.
// Takes O(N) time with O(N) space complexity.
func convertFromBase10(b10 int, b2 int) string {
	b := new(bytes.Buffer)

	for b10 > 0 {
		remainder := b10 % b2
		b10 /= b2

		if remainder >= 10 {
			b.WriteRune(decimap[remainder])
		} else {
			b.WriteRune(rune(remainder) + '0')
		}
	}

	bb := b.Bytes()

	for i := 0; i < b.Len()/2; i++ {
		bb[i], bb[b.Len()-1-i] = bb[b.Len()-1-i], bb[i]
	}

	return b.String()
}
