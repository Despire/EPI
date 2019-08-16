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

import "math"

type balancedMetadata struct {
	height   int
	balanced bool
}

// IsHeightBalanced returns true if the tree
// is height balanced (for each node the difference
// of heights for its left and right subtree is
// at most 1).
// Time complexity O(n), space complexityo O(h).
func IsHeightBalanced(root *Node) bool {
	return isBalanced(root).balanced
}

// isBalanced is a helper function that checks whether
// the tree is height balanced.
func isBalanced(root *Node) balancedMetadata {
	if root == nil {
		return balancedMetadata{
			height:   -1,
			balanced: true,
		}
	}

	left := isBalanced(root.Left)
	if !left.balanced {
		return left
	}
	right := isBalanced(root.Right)
	if !right.balanced {
		return right
	}

	return balancedMetadata{
		balanced: int(math.Abs(float64(left.height)-float64(right.height))) <= 1,
		height:   int(math.Max(float64(left.height), float64(right.height))) + 1,
	}
}

// Height returns the height (maximum depth) of the tree.
func Height(root *Node) int {
	return calcHeight(root)
}

// calcHeight is a helper function that calculates
// the hieght of the tree.
func calcHeight(root *Node) int {
	if root == nil {
		return -1
	}
	left := calcHeight(root.Left)
	right := calcHeight(root.Right)

	return 1 + int(math.Max(float64(left), float64(right)))
}
