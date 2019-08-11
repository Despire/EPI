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

// Advance advances through the array
// and returns if it is possible to
// advance to the last position.
// Time complexity O(N).
// Space Complexity O(1).
func Advance(x []uint) bool {
	max := 0
	for i := 0; i < len(x) && max >= i; i++ {
		d := i + int(x[i])
		if d > max {
			max = d
		}
	}
	return max >= len(x)-1
}
