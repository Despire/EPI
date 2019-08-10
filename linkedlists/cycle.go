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

// HasCycle checks if the linkedlists contains a cycle in
// O(N) time complexity, O(1) space complexity.
func HasCycle(root *Node) bool {
	if root == nil {
		return false
	}

	slow := root
	fast := root

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}

// RootOfCycle detects the root of the cycle if the linkedlists
// contains a cycle in O(N) time complexity, O(1) space complexity.
func RootOfCycle(root *Node) *Node {
	if !HasCycle(root) {
		return nil
	}

	slow := root
	fast := root

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			break
		}
	}

	slow = root

	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return slow
}
