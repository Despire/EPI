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

// ComputeExterior computes the exterior of a binary
// tree. Takes O(N) time with O(h) space complexity.
func ComputeExterior(root *Node) []*Node {
	if root == nil {
		return nil
	}
	result := []*Node{
		root,
	}
	result = helperLeftHalf(result, root.Left, true)
	result = helperRightHalf(result, root.Right, true)

	return result
}

func helperLeftHalf(result []*Node, root *Node, add bool) []*Node {
	// left half uses pre-order
	if root != nil {
		if add || isLeaf(root) {
			result = append(result, root)
		}
		result = helperLeftHalf(result, root.Left, add)
		result = helperLeftHalf(result, root.Right, add && root.Left == nil)
	}
	return result
}

func helperRightHalf(result []*Node, root *Node, add bool) []*Node {
	// right half uses post-order
	if root != nil {
		result = helperRightHalf(result, root.Left, add && root.Right == nil)
		result = helperRightHalf(result, root.Right, add)
		if add || isLeaf(root) {
			result = append(result, root)
		}
	}
	return result
}
