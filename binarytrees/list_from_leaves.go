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

// ListOfLeaves returns a list of leaves of the binary tree.
func ListOfLeaves(root *Node) []*Node {
	return helperLeaves(nil, root)
}

func helperLeaves(leaves []*Node, root *Node) []*Node {
	if root == nil {
		return nil
	}

	if isLeaf(root) {
		leaves = append(leaves, root)
	}
	leaves = helperLeaves(leaves, root.Left)
	leaves = helperLeaves(leaves, root.Right)

	return leaves
}
