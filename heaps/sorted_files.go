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

package heaps

import "container/heap"

// MergeFiles merges k sorted sequences into a signle one.
// Takes O(n log k) time with O(k) space complexity.
// Where k is the number of sequences and n is the total
// number of elements combined.
func MergeFiles(sequences ...[]int) ([]int, error) {
	// a heap of iterators
	iters := make(IntSliceIterHeap, 0, len(sequences))
	itemCount := 0
	for i := range sequences {
		iters.Push(NewIntSliceIter(sequences[i]))
		itemCount += len(sequences[i])
	}

	heap.Init(&iters)
	result := make([]int, 0, itemCount)
	for len(iters) != 0 {
		val, err := iters[0].Next()
		if err == Done {
			_ = heap.Pop(&iters)
			continue
		}
		if err != nil {
			return nil, err
		}
		result = append(result, val)
		heap.Fix(&iters, 0)
	}

	return result, nil
}
