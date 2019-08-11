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
	"math"
)

// digitCount returns the number of digits of the number.
// Example: if the input is 10, the function returns 2 (digit count).
func digitCount(n int) int {
	if n < 0 {
		n = -n
	}

	if n == 0 {
		return 1
	}

	return int(math.Log10(float64(n))) + 1
}

// IntToString is an implementation of a integer to string
// conversion.
// Takes O(N) time with O(N) space complexity.
func IntToString(number int) string {
	prefix := ""

	if number < 0 {
		prefix = "-"
		number = -number
	}

	if number == 0 {
		return prefix + "0"
	}

	count := digitCount(number)

	digits := make([]int, 0, count)

	for number > 0 {
		digits = append(digits, number%10)
		number /= 10
	}

	b := new(bytes.Buffer)
	b.WriteString(prefix)

	for i := count - 1; i >= 0; i-- {
		b.WriteByte(byte(digits[i] + '0'))
	}

	return b.String()
}

// StringToInt is an implementation of a string to int
// conversion function.
// Takes O(N) time with O(1) space complexity.
func StringToInt(s string) int {
	prefix := ""

	switch s[0] {
	case byte('+'):
		prefix = "+"
	case byte('-'):
		prefix = "-"
	}

	result := 0
	// get rid of the prefix
	if prefix != "" {
		s = s[1:]
	}

	for i := 0; i < len(s); i++ {
		result *= 10
		result += int(s[i] - '0')
	}

	if prefix == "-" {
		result = -result
	}

	return result
}
