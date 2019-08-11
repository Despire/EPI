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

import "unicode"

// ReverseWords reverse the words in a string s.
// The string is passed as an array of runes into the function.
// Takes O(N) time with O(1) space complexity.
func ReverseWords(chars []rune) {
	// Reverse whole sentence
	Reverse(chars)

	// Reverse each whitespace seperated word
	for i := 0; i < len(chars); i++ {
		if unicode.IsSpace(chars[i]) {
			Reverse(chars[:i])
			chars = chars[i+1:]
			i = -1
		}
	}

	// Reverse the last word in the sentence
	Reverse(chars)
}

// Reverse reverse the order.
// Takes O(N) time complexity with O(1) space.
func Reverse(chars []rune) {
	for i := 0; i < len(chars)/2; i++ {
		chars[i], chars[len(chars)-1-i] = chars[len(chars)-1-i], chars[i]
	}
}
