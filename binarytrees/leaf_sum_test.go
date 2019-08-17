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

func TestFindSum(t *testing.T) {
	tests := []struct {
		root *Node
		sum  int
		want bool
	}{
		{
			root: &Node{
				Data: 2,
				Left: &Node{
					Data: 8,
					Left: &Node{
						Data:  2,
						Left:  nil,
						Right: nil,
					},
					Right: &Node{
						Data:  4,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &Node{
					Data: 2,
					Right: &Node{
						Data:  4,
						Left:  nil,
						Right: nil,
					},
					Left: &Node{
						Data: 2,
						Left: nil,
						Right: &Node{
							Data:  2,
							Left:  nil,
							Right: nil,
						},
					},
				},
			},
			sum:  14,
			want: true,
		},
		{
			root: &Node{
				Data: 2,
				Left: &Node{
					Data: 8,
					Left: &Node{
						Data:  2,
						Left:  nil,
						Right: nil,
					},
					Right: &Node{
						Data:  4,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &Node{
					Data: 2,
					Right: &Node{
						Data:  4,
						Left:  nil,
						Right: nil,
					},
					Left: &Node{
						Data: 2,
						Left: nil,
						Right: &Node{
							Data:  2,
							Left:  nil,
							Right: nil,
						},
					},
				},
			},
			sum:  5,
			want: false,
		},
		{
			root: &Node{
				Data: 2,
				Left: &Node{
					Data: 8,
					Left: &Node{
						Data:  2,
						Left:  nil,
						Right: nil,
					},
					Right: &Node{
						Data:  4,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &Node{
					Data: 2,
					Right: &Node{
						Data:  4,
						Left:  nil,
						Right: nil,
					},
					Left: &Node{
						Data: 2,
						Left: nil,
						Right: &Node{
							Data:  2,
							Left:  nil,
							Right: nil,
						},
					},
				},
			},
			sum:  12,
			want: true,
		},
	}

	for _, test := range tests {
		got := FindSum(test.root, test.sum)

		if got != test.want {
			t.Errorf("result mismatch. got:%v; want:%v", got, test.want)
		}
	}
}
