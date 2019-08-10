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

// IsPalindrome checks if the linkedlist is a palindrom
// In O(N) time complexity, O(1) space complexity.
func IsPalindrome(root *Node) bool {
	if root == nil {
		return false
	}

	slow := root
	fast := root

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	iterator := Reverse(slow.Next)

	// the list length is 2 or less.
	if slow.Next == nil {
		return false
	}

	for root != slow {
		if root.Data.(int) != iterator.Data.(int) {
			return false
		}

		root = root.Next
		iterator = iterator.Next
	}

	return true
}
