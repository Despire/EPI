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
// limitations under the License

package stacksandqueues

import "errors"

// IntCircularQueue implements the standart queue operations.
// The implementation is a circular queue (fixed sized).
type IntCircularQueue struct {
	elems     []int
	maxSize   int
	elemCount int
	beg       int
	end       int
}

// NewCircularQueue returns an instance of type <*CircularQueue>
func NewCircularQueue(maxSize int, typ interface{}) *IntCircularQueue {
	q := &IntCircularQueue{
		elems:     make([]int, maxSize),
		maxSize:   maxSize,
		beg:       0,
		end:       0,
		elemCount: 0,
	}
	return q
}

// Enqueue inserts a new value into the queue.
func (q *IntCircularQueue) Enqueue(val int) {
	if q.elemCount == q.maxSize {
		tempSlice := make([]int, len(q.elems)*2)
		// copy elems from beg to the end of the slice.
		l := copy(tempSlice, q.elems[q.beg:])
		// copy elems that preceded beg.
		copy(tempSlice[l:], q.elems[:q.beg])

		q.beg = 0
		q.end = q.Size()

		q.elems = tempSlice
	}

	q.elems[q.end] = val
	q.end = (q.end + 1) % q.maxSize
	q.elemCount++
}

// Dequeue removes the first element in the queue.
func (q *IntCircularQueue) Dequeue() (int, error) {
	if q.elemCount == 0 {
		return 0, errors.New("empty queue")
	}
	result := q.elems[q.beg]
	q.beg = (q.beg + 1) % q.maxSize
	q.elemCount--
	return result, nil
}

// Size returns the number of elements in the queue.
func (q *IntCircularQueue) Size() int {
	return q.elemCount
}
