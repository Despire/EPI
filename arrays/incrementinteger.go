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

package arrays

// Increment increments the number passed in as a slice of digits
// by one, and returns the new number.
// Time complexity O(n).
func Increment(digits []int) []int {
	last := len(digits) - 1
	digits[last] += 1
	for i := last; i > 0 && digits[i] == 10; i-- {
		digits[i] = 0
		digits[i-1] += 1
	}
	if digits[0] == 10 {
		digits = append([]int{1, 0}, digits[1:]...)
	}
	return digits
}
