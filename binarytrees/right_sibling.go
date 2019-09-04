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

package binarytrees

type rightSiblingNode struct {
	Data       interface{}
	Left       *rightSiblingNode
	Right      *rightSiblingNode
	level_next *rightSiblingNode
}

// ComputeRightSiblingTree sets each node's level-next field
// to the node on its right if one exists. Take O(N) time with O(1) space complexity.
func ComputeRightSiblingTree(root *rightSiblingNode) {
	// Perfect binary tree has all leaves at the same depth
	// and each non-leaf has two children...

	// Outer loop traverses all the left nodes (start node of each level).
	for ;root != nil && root.Left != nil; root = root.Left {

		// traverse the whole level from left -> right.
		// and at the same time for each node set the level_next field.
		for curr := root; curr != nil; curr = curr.level_next {
			curr.Left.level_next = curr.Right

			if curr.level_next != nil {
				curr.Right.level_next = curr.level_next.Left
			}
		}
	}
}
