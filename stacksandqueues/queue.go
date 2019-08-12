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

// queue is a simply Queue data structure implementation
// using go slices.
type queue []interface{}

// Size returns the number of elements in the queue.
func (q *queue) Size() int {
	return len(*q)
}

// Enqueue inserts a new element to the end of the queue.
func (q *queue) Enqueue(val interface{}) {
	*q = append(*q, val)
}

// Dequeueu remove the first element from the queue and returns it.
func (q *queue) Dequeue() interface{} {
	result := (*q)[0]
	*q = (*q)[1:]
	return result
}
