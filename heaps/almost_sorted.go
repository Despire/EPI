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

import (
	"container/heap"
	"fmt"
)

// PrintKSorted prints almost sorted array in sorted order
// where each elements is at most k position away from it's
// correct position.
// Takes O(n log k) time with O(k) space complexity.
func PrintKSorted(iter IntSliceIter, k int) {
	// we know that if we read the k+1 value
	// the smallest value globally is present in the first k.
	// the idea is to build a heap o k items and if the sequence contains another
	// value the root of the heap will be the current smallest element.
	kSmallest := make(IntHeap, 0, k)

	for i := 0; i < k; i++ {
		val, err := iter.Next()
		if err == Done {
			break
		}
		heap.Push(&kSmallest, val)
	}

	for {
		val, err := iter.Next()
		if err == Done {
			break
		}
		fmt.Println(heap.Pop(&kSmallest).(int))
		heap.Push(&kSmallest, val)
	}

	for len(kSmallest) > 0 {
		fmt.Println(heap.Pop(&kSmallest).(int))
	}
}
