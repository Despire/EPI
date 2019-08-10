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

package linkedlists

// Reverse reverses a linked lists in O(N) time complexity.
func Reverse(root *Node) *Node {

	prev := (*Node)(nil)
	curr := root

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

// ReverseSubList reverses a sublist of a linked list
// in O(N) time complexity, O(1) space complexity.
func ReverseSubList(root *Node, i, j int) *Node {
	dummy := &Node{
		Data: 0,
		Next: root,
	}

	head := dummy
	for x := 1; x < i; x++ {
		head = head.Next
	}

	iter := head.Next
	for i < j {
		next := iter.Next
		iter.Next = next.Next
		next.Next = head.Next
		head.Next = next
		i++
	}
	return dummy.Next
}
