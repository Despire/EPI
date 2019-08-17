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

// IterativeInorder returns a pre-order traversal
// for the given binary tree.
// Takes O(N) time with O(h) space complexity.
func IterativeInorder(root *Node) []*Node {
	pop := func(s *[]*Node) *Node {
		r := (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]
		return r
	}
	push := func(s *[]*Node, val *Node) {
		*s = append(*s, val)
	}
	var stack, traverse []*Node

	curr := root
	for curr != nil {
		if curr != nil {
			push(&stack, curr)
			curr = curr.Left
		} else {
			curr = pop(&stack)
			traverse = append(traverse, curr)
			curr = curr.Right
		}
	}

	return traverse
}
