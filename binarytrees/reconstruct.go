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

// Reconstruct reconstruts a binary tree using the given
// preorder and inorder traversals. Returns the root of the
// binarytree.
func Reconstruct(preorder, inorder []string) *Node {
	nameToIdx := map[string]int{}
	for k, v := range inorder {
		nameToIdx[v] = k
	}
	return helperReconstruct(preorder, 0, len(preorder), 0, len(inorder), nameToIdx)
}

func helperReconstruct(preorder []string, preStart, preEnd, inStart, inEnd int,nameToIdx map[string]int) *Node {
	if preEnd <= preStart || inEnd <= inStart {
		return nil
	}

	RootIdx := nameToIdx[preorder[preStart]]
	LeftSize := RootIdx - inStart

	return &Node{
		Data: preorder[preStart],
		Left: helperReconstruct(preorder, preStart + 1, preStart + 1 + LeftSize, inStart, RootIdx, nameToIdx),
		Right: helperReconstruct(preorder, preStart + 1 + LeftSize, preEnd, RootIdx + 1, inEnd, nameToIdx),
	}
}

