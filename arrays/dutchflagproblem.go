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

// Partitions the slice in linear time, such that all elements
// less than the pivot appears before all elements
// equal to the pivot and elements which are greater than the pivot
// appears as last.
// The pivot is selected as elems[i].
func Partition(elems []int, i int) (int, int) {
	// paritions the slice in O(n) time.
	// such that elems[0:s] contains all elements that are '<p'
	// elems[s:g] contains all elements that are '=p'
	// elems[g:] contains all the elements that are '>p"
	p := elems[i]
	s := 0
	g := len(elems)
	e := 0

	for e < g {
		switch {
		case elems[e] < p:
			elems[s], elems[e] = elems[e], elems[s]
			s += 1
			e += 1
		case elems[e] > p:
			g -= 1
			elems[g], elems[e] = elems[e], elems[g]
		default:
			// if isn't '>p' '<p' than it is equal
			e += 1
		}
	}
	return s, g
}
