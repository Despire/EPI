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

type helper struct {
	currNode int
	node     *Node
}

// KthNodeInorderTraversal returns the Kth node in an inorder
// traversal. Takes O(N) time with O(h) space complexity.
func KthNodeInorderTraversal(k int, root *Node) *Node {
	result := &helper{
		currNode: 0,
		node:     nil,
	}
	helperFunc(k, root, result)
	return result.node
}

func helperFunc(k int, root *Node, h *helper) {
	if root == nil {
		return
	}
	helperFunc(k, root.Left, h)
	if h.currNode == k {
		return
	}
	h.currNode += 1
	if h.currNode == k {
		h.node = root
		return
	}
	helperFunc(k, root.Right, h)
}

type ModifiedNode struct {
	Data        interface{}
	Left        *ModifiedNode
	Right       *ModifiedNode
	subtreeSize int
}

// KthNodeInorderTraversalWithModifiedNode does the same thing
// as KthNodeInorderTraversal but uses a modified binary tree
// node wich contains an addional field that hold the number
// of nodes for the subtree rooted at that node.
func KthNodeInorderTraversalWithModifiedNode(k int, root *ModifiedNode) *ModifiedNode {
	iter := root
	for iter != nil {
		leftTreeSize := 0
		if iter.Left != nil {
			leftTreeSize = iter.Left.subtreeSize
		}
		if leftTreeSize >= k {
			iter = iter.Left
		} else {
			k -= leftTreeSize
			// its not the left or right child
			if k == 1 {
				break
			}
			// if is not the root then
			// the node must be in the
			// right subtree.
			k -= 1
			iter = iter.Right
		}
	}
	return iter
}
