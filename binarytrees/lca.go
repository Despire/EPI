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

// exists returns true if <node> is a node
// in the tree that starts at <root>
// Take O(N) time with O(h) space complexity.
func exists(node *Node, root *Node) bool {
	if root == nil {
		return false
	}
	return root == node || exists(node, root.Left) || exists(node, root.Right)
}

func lcahelper(root, node0, node1 *Node) (*Node, int) {
	if root == nil {
		return nil, 0
	}
	lt, cl := lcahelper(root.Left, node0, node1)
	if cl == 2 {
		return lt, cl
	}
	rt, cr := lcahelper(root.Right, node0, node1)
	if cr == 2 {
		return rt, cr
	}

	count := cl + cr

	if root == node0 {
		count++
	}
	if root == node1 {
		count++
	}

	if count == 2 {
		return root, count
	}
	return nil, count
}

// LCA finds the least common ancestor of two nodes
// in a binary tree. Returns nil if no LCA is found.
// Take O(N) time with O(h) space complexity.
func LCA(root, node0, node1 *Node) *Node {
	r, _ := lcahelper(root, node0, node1)
	return r
}
