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

// ApplyPermutation will apply the permutation
// in O(N) time and with additional O(N) space complexity.
func ApplyPermutation(elems, permutation []int) {
	s := make([]bool, len(elems))

	for i := range elems {
		current := i
		for s[i] == false {
			elems[i], elems[permutation[current]] = elems[permutation[current]], elems[i]
			s[permutation[current]] = true
			current = permutation[current]
		}
	}

}
