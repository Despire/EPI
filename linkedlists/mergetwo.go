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

// Merge merges two sorted linked lists in O(N) time complexity.
// O(1) space complexity.
func Merge(first *Node, second *Node) *Node {
	var root = &Node{
		Data: 0,
		Next: nil,
	}

	var fwd = root
	for ; first != nil && second != nil; fwd = fwd.Next {
		if first.Data.(int) < second.Data.(int) {
			fwd.Next = first
			first = first.Next
		} else {
			fwd.Next = second
			second = second.Next
		}
	}

	if first != nil {
		fwd.Next = first
	} else {
		fwd.Next = second
	}

	return root.Next
}
