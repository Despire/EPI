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

import "testing"

func TestListOfLeaves(t *testing.T) {
	tree1 := &Node{
		Data: 1,
		Left: &Node{
			Data: 2,
			Left: &Node{
				Data: 4,
			},
			Right: &Node{
				Data: 5,
			},
		},
		Right: &Node{
			Data: 3,
			Left: &Node{
				Data: 6,
			},
		},
	}

	tests := []struct {
		root *Node
		want []int
	}{
		{
			root: tree1,
			want: []int{4, 5, 6},
		},
		{
			root: nil,
			want: nil,
		},
		{
			root: &Node{Data: 5},
			want: []int{5},
		},
	}

	for _, test := range tests {
		have := ListOfLeaves(test.root)
		for i, _ := range have {
			if have[i].Data.(int) != test.want[i] {
				t.Errorf("result mismatch! have: %v; want: %v", have[i].Data.(int), test.want[i])
			}
		}
	}
}
