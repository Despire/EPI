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

func TestIsSymetric(t *testing.T) {
	tests := []struct {
		root *Node
		want bool
	}{
		{
			root: &Node{
				Data: 314,
				Left: &Node{
					Data: 6,
					Left: nil,
					Right: &Node{
						Data: 2,
						Left: nil,
						Right: &Node{
							Data:  3,
							Left:  nil,
							Right: nil,
						},
					},
				},
				Right: &Node{
					Data: 6,
					Left: &Node{
						Data: 2,
						Left: &Node{
							Data:  3,
							Left:  nil,
							Right: nil,
						},
						Right: nil,
					},
					Right: nil,
				},
			},
			want: true,
		},
		{
			root: &Node{
				Data: 314,
				Left: &Node{
					Data:  5,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
			want: false,
		},

		{
			root: &Node{
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
			},
			want: false,
		},
		{
			root: nil,
			want: false,
		},
		{
			root: &Node{
				Data: 5,
			},
			want: true,
		},
	}
	for _, test := range tests {
		got := IsSymetric(test.root)

		if got != test.want {
			t.Errorf("result mismatch. got: %v; want: %v", got, test.want)
		}
	}
}
