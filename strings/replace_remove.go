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

// ReplaceAndRemove replaces each 'a' by two 'd' and removes
// each occurence of 'b'
// Takes O(N) time with O(1) space complexity.
func ReplaceAndRemove(chars []rune, size int) int {
	aCount := 0
	read := 0
	write := 0

	// remove "b's" and count the number of "a's".
	for read < size {
		if chars[read] == 'a' {
			aCount++
		}
		if chars[read] != 'b' {
			chars[write] = chars[read]
			write++
		}
		read++
	}

	// replace "a's" with "dd's" starting from the end.
	read = write - 1
	write = write + aCount - 1
	size = write + 1
	for aCount > 0 {
		if chars[read] == 'a' {
			aCount--

			chars[write] = 'd'
			chars[write-1] = 'd'

			write--
		} else {
			chars[write] = chars[read]
		}
		read--
		write--
	}

	return size
}
