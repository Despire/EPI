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

// IsSymetric returns <true> if the tree is symetric.
// Takes O(N) time and O(h) space complexity.
func IsSymetric(root *Node) bool {
	if root == nil {
		return false
	}
	return isSymetric(root.Left, root.Right)
}

func isSymetric(left *Node, right *Node) bool {
	if left == nil && right == nil {
		return true
	}
	if left != nil && right != nil {
		return left.Data == right.Data &&
			isSymetric(left.Right, right.Left) &&
			isSymetric(left.Left, right.Right)
	}
	return false
}
