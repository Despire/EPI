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

func TestComputeSum(t *testing.T) {
	tests := []struct {
		root *Node
		want int
	}{
		{
			root: &Node{
				Data: 1,
				Left: &Node{
					Data: 0,
					Left: &Node{
						Data: 0,
						Left: &Node{
							Data:  0,
							Left:  nil,
							Right: nil,
						},
						Right: &Node{
							Data:  1,
							Left:  nil,
							Right: nil,
						},
					},
					Right: &Node{
						Data: 1,
						Left: nil,
						Right: &Node{
							Data: 1,
							Left: &Node{
								Data:  0,
								Left:  nil,
								Right: nil,
							},
							Right: nil,
						},
					},
				},
				Right: &Node{
					Data: 1,
					Left: &Node{
						Data: 0,
						Left: nil,
						Right: &Node{
							Data: 0,
							Left: &Node{
								Data: 1,
								Left: nil,
								Right: &Node{
									Data:  1,
									Left:  nil,
									Right: nil,
								},
							},
							Right: &Node{
								Data:  0,
								Left:  nil,
								Right: nil,
							},
						},
					},
					Right: &Node{
						Data: 0,
						Left: nil,
						Right: &Node{
							Data:  0,
							Left:  nil,
							Right: nil,
						},
					},
				},
			},
			want: 126,
		},
		{
			root: &Node{
				Data: 1,
				Left: &Node{
					Data: 0,
					Left: nil,
					Right: &Node{
						Data:  1,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &Node{
					Data:  0,
					Left:  nil,
					Right: nil,
				},
			},
			want: 7,
		},
	}

	for _, test := range tests {
		got := ComputeSum(test.root)
		if got != test.want {
			t.Errorf("result mismatch. got: %v; want: %v", got, test.want)
		}
	}
}
