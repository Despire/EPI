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
	"strconv"
	"strings"
	"unicode"
)

// RLE (acronym for Run length encoding) is an compression algorithm
// of strings.
// Takes O(N) time complexity.
func RLE(s string) string {
	buffer := new(bytes.Buffer)
	for i := 0; i < len(s); i++ {
		count := 1
		for i+1 < len(s) && s[i] == s[i+1] {
			i++
			count++
		}
		buffer.WriteString(strconv.Itoa(count))
		buffer.WriteByte(s[i])
	}

	return buffer.String()
}

// RLD (acronym for Run length decoding) is an decoding function
// for a string encoded using RLE.
// Takes O(N) time complexity.
func RLD(s string) string {
	buffer := new(strings.Builder)

	for i := 0; i < len(s); i++ {
		count := 0
		for unicode.IsDigit(rune(s[i])) && i < len(s) {
			count = count*10 + int(s[i]-'0')
			i++
		}
		char := s[i : i+1]
		for j := 0; j < count; j++ {
			buffer.WriteString(char)
		}
	}

	return buffer.String()
}
