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

// NodeWithParent is a modified binary tree
// node, in addition to the basic fields
// it also contains a pointer to the parent node.
type NodeWithParent struct {
	Data   interface{}
	Left   *NodeWithParent
	Right  *NodeWithParent
	Parent *NodeWithParent
}

// Next returns the successor of the node.
// The sucessor is returned in an inorder fashion.
func Next(node *NodeWithParent) *NodeWithParent {
	iter := node
	if iter == nil {
		return nil
	}

	if iter.Right != nil {
		iter = iter.Right
		for iter.Left != nil {
			iter = iter.Left
		}
		return iter
	}

	for iter.Parent != nil && iter == iter.Parent.Right {
		iter = iter.Parent
	}

	return iter.Parent
}
