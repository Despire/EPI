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

// QueueStack is a queue implemented using stacks.
type QueueStack struct {
	elems *stack
	deque *stack
}

// Enqueue inserts a new elements into the queue.
func (q *QueueStack) Enqueue(val int) {
	q.elems.Push(val)
}

// Dequeue removes an element from the queue
// and returns it value.
// Takes O(m) time complexity.
func (q *QueueStack) Dequeue() int {
	if q.deque.Size() == 0 {
		for q.elems.Size() > 0 {
			q.deque.Push(q.elems.Pop())
		}
	}
	return q.deque.Pop().(int)
}
