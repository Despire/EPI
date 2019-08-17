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

// FindSum finds a root-to-leaf path with the given sum
// Takes O(N) time with O(h) space complexity.
func FindSum(root *Node, sum int) bool {
	return findSumHelper(root, sum, 0)
}

func findSumHelper(root *Node, sum, curr int) bool {
	if root == nil {
		return false
	}
	curr += root.Data.(int)
	return (isLeaf(root) && sum == curr) || findSumHelper(root.Left, sum, curr) || findSumHelper(root.Right, sum, curr)
}
