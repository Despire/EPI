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

import "sort"

// Alternation rearranges the slice such that
// the elements are in the order:
// x[0] <= x[1] >= x[2] <= x[3] >= x[4] ...
// Takes O(nlogn) time and O(1) space
func Alternation(x []int) {
	sort.Ints(x)
	for i := 2; i < len(x); i += 2 {
		x[i-1], x[i] = x[i], x[i-1]
	}
}

// LinearAlternation rearranges the slice such that
// the elements are in the order:
// x[0] <= x[1] >= x[2] <= x[3] >= x[4] ...
// Takes O(n) time and O(1) space
func LinearAlternation(x []int) {
	for i := 1; i < len(x); i += 1 {
		switch {
		case i%2 == 0:
			if x[i-1] < x[i] {
				x[i], x[i-1] = x[i-1], x[i]
			}
		case i%2 == 1:
			if x[i-1] > x[i] {
				x[i], x[i-1] = x[i-1], x[i]
			}
		}
	}
}
