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

import (
	"testing"
)

func TestNext(t *testing.T) {
	root := &NodeWithParent{
		Data: 314,
		Left: &NodeWithParent{
			Data: 6,
			Left: &NodeWithParent{
				Data: 271,
				Left: &NodeWithParent{
					Data: 28,
				},
				Right: &NodeWithParent{
					Data: 0,
				},
			},
			Right: &NodeWithParent{
				Data: 561,
				Right: &NodeWithParent{
					Data: 3,
					Left: &NodeWithParent{
						Data: 17,
					},
				},
			},
		},
		Right: &NodeWithParent{
			Data: 6,
			Left: &NodeWithParent{
				Data: 2,
				Right: &NodeWithParent{
					Data: 1,
					Left: &NodeWithParent{
						Data: 401,
						Right: &NodeWithParent{
							Data: 641,
						},
					},
					Right: &NodeWithParent{
						Data: 257,
					},
				},
			},
			Right: &NodeWithParent{
				Data: 271,
				Right: &NodeWithParent{
					Data: 28,
				},
			},
		},
	}

	root.Left.Parent = root
	root.Right.Parent = root

	root.Left.Left.Parent = root.Left
	root.Left.Right.Parent = root.Left

	root.Left.Left.Left.Parent = root.Left.Left
	root.Left.Left.Right.Parent = root.Left.Left

	root.Left.Right.Right.Parent = root.Left.Right
	root.Left.Right.Right.Left.Parent = root.Left.Right.Right

	root.Right.Right.Parent = root.Right
	root.Right.Right.Right.Parent = root.Right.Right

	root.Right.Left.Parent = root.Right
	root.Right.Left.Right.Parent = root.Right.Left
	root.Right.Left.Right.Right.Parent = root.Right.Left.Right
	root.Right.Left.Right.Left.Parent = root.Right.Left.Right
	root.Right.Left.Right.Left.Right.Parent = root.Right.Left.Right.Left

	tests := []struct {
		node *NodeWithParent
		want *NodeWithParent
	}{
		{
			node: root.Left.Right,
			want: root.Left.Right.Right.Left,
		},
		{
			node: root.Left.Right.Right.Left,
			want: root.Left.Right.Right,
		},
		{
			node: root.Left.Right.Right,
			want: root,
		},
		{
			node: nil,
			want: nil,
		},
		{
			node: root.Right.Right.Right,
			want: nil,
		},
		{
			node: root.Left.Left.Right,
			want: root.Left,
		},
	}

	for _, test := range tests {
		have := Next(test.node)
		if have != test.want {
			t.Errorf("result mismatch. have: %v; want: %v", have, test.want)
		}
	}
}
