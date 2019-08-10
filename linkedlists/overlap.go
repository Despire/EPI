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

// DoOverlap checks if two linkedlists overlap
// In O(N) time complexity.
func DoOverlap(root1 *Node, root2 *Node) bool {
	root1Size := 0
	root2Size := 0

	dummy1 := root1
	for dummy1 != nil {
		dummy1 = dummy1.Next
		root1Size++
	}

	dumm2 := root2
	for dumm2 != nil {
		dumm2 = dumm2.Next
		root2Size++
	}

	if root1Size > root2Size {
		root1 = advanceByK(root1, root1Size-root2Size)
	} else {
		root2 = advanceByK(root2, root2Size-root1Size)
	}

	for root1 != nil && root2 != nil {
		root1 = root1.Next
		root2 = root2.Next

		if root1 == root2 {
			return true
		}
	}
	return false
}

// advanceByK returns the kth node starting from <root> node.
func advanceByK(root *Node, k int) *Node {
	for i := 0; i < k; i++ {
		root = root.Next
	}

	return root
}
