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

// Rotate2DArray takes as input an nxn 2D array, and rotates
// the array by 90 degrees clockwise.
// Takes O(N^2) time and O(1) space complexity.
func Rotate2DArray(matrix [][]int) {
	rows := len(matrix)

	for offset := 0; offset < rows/2; offset++ {
		for i := offset; i < rows-1-offset; i++ {
			// top left with top right
			matrix[offset][i], matrix[i][rows-1-offset] = matrix[i][rows-1-offset], matrix[offset][i]

			// top left with bottom right
			matrix[offset][i], matrix[rows-1-offset][rows-1-i] = matrix[rows-1-offset][rows-1-i], matrix[offset][i]

			// top left with bottom left
			matrix[offset][i], matrix[rows-1-i][offset] = matrix[rows-1-i][offset], matrix[offset][i]
		}
	}
}
