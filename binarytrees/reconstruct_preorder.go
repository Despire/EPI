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

func ReconstructFromPreorder(preorder []string) *Node {
	root, _ := helperReconstructFromPreorder(preorder, 0)
	return root
}

func helperReconstructFromPreorder(preorder []string, curr int) (*Node, int) {
	if len(preorder) == 0 || preorder[curr] == "null" {
		return nil, curr
	}
	key := preorder[curr]

	left, curr := helperReconstructFromPreorder(preorder, curr + 1)
	right, curr := helperReconstructFromPreorder(preorder, curr + 1)

	return &Node{
		Data: key,
		Left: left,
		Right: right,
	}, curr
}
