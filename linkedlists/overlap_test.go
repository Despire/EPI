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

package linkedlists

import "testing"

func TestDoOverlap(t *testing.T) {
	tests := []struct {
		inputFirst  *Node
		inputSecond *Node
		want        bool
	}{
		{
			inputFirst: &Node{
				Data: 5,
				Next: &Node{
					Data: 10,
					Next: &Node{
						Data: 20,
						Next: &Node{
							Data: 50,
							Next: &Node{
								Data: 60,
								Next: &Node{
									Data: 70,
									Next: nil,
								},
							},
						},
					},
				},
			},

			inputSecond: &Node{
				Data: 1,
				Next: &Node{
					Data: 2,
					Next: &Node{
						Data: 3,
						Next: nil,
					},
				},
			},

			want: true,
		},
		{
			inputFirst: &Node{
				Data: 5,
				Next: &Node{
					Data: 10,
					Next: &Node{
						Data: 20,
						Next: &Node{
							Data: 50,
							Next: &Node{
								Data: 60,
								Next: &Node{
									Data: 70,
									Next: nil,
								},
							},
						},
					},
				},
			},

			want: false,
		},
		{
			inputFirst: &Node{
				Data: 10,
				Next: nil,
			},
			want: false,
		},
		{
			inputFirst: &Node{
				Data: 10,
				Next: &Node{
					Data: 20,
					Next: nil,
				},
			},
			want: false,
		},
	}

	tests[0].inputSecond.Next.Next.Next = tests[0].inputFirst.Next.Next.Next
	for _, test := range tests {
		got := DoOverlap(test.inputFirst, test.inputSecond)

		if got != test.want {
			t.Errorf("unwated result! got: %v; want: %v", got, test.want)
		}
	}
}
