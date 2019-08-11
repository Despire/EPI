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

import "math/rand"

// if we have a set of size 1 it's trivial just return the element
// if we have a set of 2 and subsetSize is 1 just return the random of the two.
// if we have a set of 3 and subsetSize is 1 just return the random of the three.
// if we have a set of 3 and subsetSize is 2 just return the random of three and then random of two.
// if we have a set of 3 and subsetSize is 3 just return the random of the first then random of the two and than the last element left.

// sampleOfflineData returns a subset of the given size of the array elements.
// All subsets all equal likely. Takes O(N) time with O(1) space complexity.
func sampleOfflineData(data []int, subsetSize int) []int {
	if subsetSize < 1 {
		return data[:0]
	}

	for i := range data {
		if i >= subsetSize {
			break
		}
		startElement := i + rand.Intn(len(data)-i)
		data[i], data[startElement] = data[startElement], data[i]
	}

	return data[:subsetSize]
}
