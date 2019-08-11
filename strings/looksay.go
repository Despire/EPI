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
)

// CalcNthLookSayNumber calculated the nth look and say number
// starting from the first one.
// Takes O(N*(2^n)) time complexity.
func CalcNthLookSayNumber(n int) string {
	s := "1"
	for i := 1; i < n; i++ {
		s = next(s)
	}
	return s
}

// next calculates the next look and say number from the one
// passed in as <s>.
func next(s string) string {
	buffer := new(bytes.Buffer)

	for i := 0; i < len(s); i++ {
		count := 1
		num := s[i]

		for i+1 < len(s) && s[i+1] == s[i] {
			count++
			i++
		}

		buffer.WriteByte(byte(count + '0'))
		buffer.WriteByte(num)
	}

	return buffer.String()
}
