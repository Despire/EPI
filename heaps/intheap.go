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

import "errors"

var (
	Done = errors.New("no more values left")
)

type IntSliceIter struct {
	pos   int
	slice []int
}

func (iter *IntSliceIter) Next() (int, error) {
	if iter.pos == len(iter.slice) {
		return 0, Done
	}
	result := iter.slice[iter.pos]
	iter.pos++
	return result, nil
}

func (iter *IntSliceIter) Less(other IntSliceIter) bool {
	// iter is invalid therefore it is always less.
	if iter.pos == len(iter.slice) {
		return true
	}
	// other iter is invalid therefore this one is not less.
	if other.pos == len(other.slice) {
		return false
	}
	return iter.slice[iter.pos] < other.slice[other.pos]
}

func NewIntSliceIter(s []int) IntSliceIter {
	return IntSliceIter{
		pos:   0,
		slice: s,
	}
}

// An IntSliceIterHeap is a min-heap of IntSliceIter.
type IntSliceIterHeap []IntSliceIter

func (h IntSliceIterHeap) Len() int           { return len(h) }
func (h IntSliceIterHeap) Less(i, j int) bool { return h[i].Less(h[j]) }
func (h IntSliceIterHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntSliceIterHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(IntSliceIter))
}

func (h *IntSliceIterHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
