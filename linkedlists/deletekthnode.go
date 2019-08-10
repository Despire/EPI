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

// DeleteKthLast deletes the kth last node from a linkedlists
// In O(N) time complexity, O(1) space complexity.
func DeleteKthLast(node *Node, k int) bool {
	if node == nil {
		return false
	}

	first := node
	secnd := node

	for first.Next != nil && k > 0 {
		first = first.Next
		k--
	}

	if k > 0 {
		// the list is not atleast k nodes long
		return false
	}

	for first.Next != nil {
		first = first.Next
		secnd = secnd.Next
	}

	secnd.Next = secnd.Next.Next

	return true
}
