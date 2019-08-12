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

package stacksandqueues

import (
	"github.com/despire/EPI/binarytrees"
)

// OrderOfIncreasingDepth returns an array consisting of the keys at
// the same level. Keys appear in the order of the corresponding
// nodes'depths from left to right.
// Time complexity O(N).
func OrderOfIncreasingDepth(root *binarytrees.Node) [][]interface{} {
	if root == nil {
		return nil
	}
	var elemsPerLevel [][]interface{}
	queue := new(queue)

	// level 0
	queue.Enqueue(root)

	// while there are still levels to traverse.
	for queue.Size() > 0 {
		// traverse all nodes in the current level.
		var currLevel []interface{}
		for size := queue.Size(); size > 0; size-- {
			curr := queue.Dequeue().(*binarytrees.Node)
			currLevel = append(currLevel, curr.Data)
			if curr.Left != nil {
				queue.Enqueue(curr.Left)
			}
			if curr.Right != nil {
				queue.Enqueue(curr.Right)
			}
		}
		elemsPerLevel = append(elemsPerLevel, currLevel)
	}

	return elemsPerLevel
}
