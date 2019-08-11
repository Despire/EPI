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

import "math"

// SpiralOrdering takes an nxn 2D array and returns
// the spral ordering of the array.
// Takes O(N^2) time and O(1) space complexity.
func SpiralOrdering(matrix [][]int) []int {
	rows := len(matrix)
	cols := len(matrix[0])

	result := make([]int, 0)
	for offset := 0; offset < int(math.Ceil(float64(rows)/2.0)); offset++ {
		if offset == cols-1-offset {
			result = append(result, matrix[offset][offset])
			continue
		}
		for i := offset; i < cols-1-offset; i++ {
			result = append(result, matrix[offset][i])
		}
		for i := offset; i < rows-1-offset; i++ {
			result = append(result, matrix[i][cols-1-offset])
		}
		for i := cols - 1 - offset; i > offset; i-- {
			result = append(result, matrix[rows-1-offset][i])
		}
		for i := rows - 1 - offset; i > offset; i-- {
			result = append(result, matrix[i][offset])
		}
	}
	return result
}
