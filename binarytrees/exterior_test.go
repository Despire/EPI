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
	"reflect"
	"testing"
)

func TestComputeExterior(t *testing.T) {
	tree1 := &Node{
		Data: 1,
		Left: &Node{
			Data: 2,
		},
		Right: &Node{
			Data: 3,
		},
	}

	tree2 := &Node{
		Data: 314,
		Left: &Node{
			Data: 6,
			Left: &Node{
				Data: 271,
				Left: &Node{
					Data:  28,
					Left:  nil,
					Right: nil,
				},
				Right: &Node{
					Data:  0,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &Node{
				Data: 561,
				Left: nil,
				Right: &Node{
					Data: 3,
					Left: &Node{
						Data:  17,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
			},
		},
		Right: &Node{
			Data: 6,
			Left: &Node{
				Data: 2,
				Left: nil,
				Right: &Node{
					Data: 1,
					Left: &Node{
						Data: 401,
						Left: nil,
						Right: &Node{
							Data:  641,
							Left:  nil,
							Right: nil,
						},
					},
					Right: &Node{
						Data:  257,
						Left:  nil,
						Right: nil,
					},
				},
			},
			Right: &Node{
				Data: 271,
				Left: nil,
				Right: &Node{
					Data:  28,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}

	tree3 := &Node{
		Data: 5,
		Left: &Node{
			Data: 3,
			Right: &Node{
				Data: 5,
				Left: &Node{
					Data: 10,
				},
			},
		},
		Right: &Node{
			Data: 15,
		},
	}

	tree4 := &Node{
		Data: 3,
		Left: &Node{
			Data: 2,
			Left: &Node{
				Data: 1,
				Right: &Node{
					Data: -1,
				},
			},
			Right: &Node{
				Data: 0,
				Right: &Node{
					Data: -2,
				},
			},
		},
		Right: &Node{
			Data: 5,
			Left: &Node{
				Data: 4,
			},
			Right: &Node{
				Data: 6,
			},
		},
	}

	tests := []struct {
		tree *Node
		want []*Node
	}{
		{
			tree: tree1,
			want: []*Node{
				tree1,
				tree1.Left,
				tree1.Right,
			},
		},
		{
			tree: tree2,
			want: []*Node{
				tree2,
				tree2.Left,
				tree2.Left.Left,
				tree2.Left.Left.Left,
				tree2.Left.Left.Right,
				tree2.Left.Right.Right.Left,
				tree2.Right.Left.Right.Left.Right,
				tree2.Right.Left.Right.Right,
				tree2.Right.Right.Right,
				tree2.Right.Right,
				tree2.Right,
			},
		},
		{
			tree: tree3,
			want: []*Node{
				tree3,
				tree3.Left,
				tree3.Left.Right,
				tree3.Left.Right.Left,
				tree3.Right,
			},
		},
		{
			tree: tree4,
			want: []*Node{
				tree4,
				tree4.Left,
				tree4.Left.Left,
				tree4.Left.Left.Right,
				tree4.Left.Right.Right,
				tree4.Right.Left,
				tree4.Right.Right,
				tree4.Right,
			},
		},
	}

	for _, test := range tests {
		have := ComputeExterior(test.tree)

		if !reflect.DeepEqual(test.want, have) {
			t.Errorf("result mismatch! have:%v\n want:%v", have, test.want)
		}
	}
}
