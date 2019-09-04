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

// Rules...
// 1: a node's state cannot be set to lock if any of its descendant are in lock.
// 2: a node's state cannot be set to lock if any of its ancestors are in lock.
type NodeWithLock struct {
	Parent *NodeWithLock
	Left   *NodeWithLock
	Right  *NodeWithLock

	isLocked bool
	descendantsLocked int
}

// IsLocked returns true if the node's state is locked.
func (n *NodeWithLock) IsLocked() bool {
	return n.isLocked
}

// Lock returns true if the node's state was locked successfully.
func (n *NodeWithLock) Lock() bool {
	if n.descendantsLocked > 0 || n.isLocked {
		return false
	}

	for curr := n.Parent; curr != nil; curr = curr.Parent {
		if curr.isLocked {
			return false
		}
	}

	n.isLocked = true

	for curr := n.Parent; curr != nil; curr = curr.Parent {
		curr.descendantsLocked++
	}

	return true
}

// Unlock unlocks the nodes state.
func (n *NodeWithLock) Unlock() bool {
	if n.isLocked {
		n.isLocked = false
		for curr := n.Parent; curr != nil; curr = curr.Parent {
			curr.descendantsLocked--
		}
	}
}
