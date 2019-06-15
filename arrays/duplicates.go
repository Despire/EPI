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

// DeleteDuplicates deletes duplicate elements
// and returns the number of valid elements left
// in the slice.
// Takes O(n) time and O(1) space
func DeleteDuplicates(x []int) int {
	if len(x) == 0 {
		return 0
	}
	last := 0
	for current := 1; current < len(x); current++ {
		if x[current] != x[last] {
			last += 1
			x[last] = x[current]
		}
	}
	return last + 1
}
