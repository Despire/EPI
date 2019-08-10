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

// ShiftRightByK shifts the linked list by K-time
// in O(N) time complexity, O(1) space complexity.
func ShiftRightByK(root *Node, k int) *Node {
	if root == nil {
		return nil
	}

	first := root
	secnd := root

	for i := 0; i < k-1 && first != nil; i++ {
		first = first.Next
	}

	if first == nil {
		return nil
	}

	for first.Next != nil {
		first = first.Next
		secnd = secnd.Next
	}

	newRoot := secnd
	first.Next = root

	for root.Next != newRoot {
		root = root.Next
	}

	root.Next = nil

	return newRoot

}
