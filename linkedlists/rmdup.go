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

// RemoveDuplicates removes duplicate nodes
// from a linked list in O(N) time complexity
// O(1) space complexity.
func RemoveDuplicates(root *Node) {
	if root == nil {
		return
	}

	start := root
	end := root.Next

	for end != nil {
		if start.Data.(int) == end.Data.(int) {
			end = end.Next
		} else {
			start.Next = end
			start = end
			end = end.Next
		}
	}

	start.Next = end
}
