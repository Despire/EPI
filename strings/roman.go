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
// limitations under the License

package strings

// symbols must appear in nonincreasing order
// romanToDeci maps roman characters to their numeric values.
var romanToDeci = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

// RomanToDecimal converts a roman integer passed in as a string
// and returns it's value in base 10.
// Time complexity O(N) with O(1) space.
func RomanToDecimal(s string) int {
	sum := romanToDeci[s[len(s)-1]]

	for i := len(s) - 2; i >= 0; i-- {
		if romanToDeci[s[i]] > romanToDeci[s[i+1]] {
			sum += romanToDeci[s[i]]
		} else {
			sum -= romanToDeci[s[i]]
		}
	}
	return sum
}
