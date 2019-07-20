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

import (
	"reflect"
	"testing"
)

func TestEvenOddMerge(t *testing.T) {
	tests := []struct {
		input *Node
		want  []int
	}{
		{
			input: &Node{
				Data: 1,
				Next: &Node{
					Data: 2,
					Next: &Node{
						Data: 3,
						Next: &Node{
							Data: 4,
							Next: nil,
						},
					},
				},
			},
			want: []int{1, 3, 2, 4},
		},
		{
			input: &Node{
				Data: 1,
				Next: &Node{
					Data: 2,
					Next: &Node{
						Data: 3,
						Next: &Node{
							Data: 4,
							Next: &Node{
								Data: 5,
								Next: nil,
							},
						},
					},
				},
			},
			want: []int{1, 3, 5, 2, 4},
		},
		{
			input: &Node{
				Data: 1,
				Next: &Node{
					Data: 2,
					Next: &Node{
						Data: 3,
						Next: &Node{
							Data: 4,
							Next: &Node{
								Data: 5,
								Next: &Node{
									Data: 6,
									Next: nil,
								},
							},
						},
					},
				},
			},
			want: []int{1, 3, 5, 2, 4, 6},
		},
		{
			input: &Node{
				Data: 1,
				Next: &Node{
					Data: 2,
					Next: &Node{
						Data: 3,
						Next: &Node{
							Data: 4,
							Next: &Node{
								Data: 5,
								Next: &Node{
									Data: 6,
									Next: &Node{
										Data: 7,
										Next: nil,
									},
								},
							},
						},
					},
				},
			},
			want: []int{1, 3, 5, 7, 2, 4, 6},
		},
		{
			input: &Node{
				Data: 1,
				Next: &Node{
					Data: 2,
					Next: nil,
				},
			},
			want: []int{1, 2},
		},
		{
			input: &Node{
				Data: 1,
			},
			want: []int{1},
		},
		{
			input: nil,
			want:  []int{},
		},
	}

	for _, test := range tests {
		gotRoot := EvenOddMerge(test.input)

		gotSlice := []int{}

		for gotRoot != nil {
			gotSlice = append(gotSlice, gotRoot.Data.(int))
			gotRoot = gotRoot.Next
		}

		if !reflect.DeepEqual(gotSlice, test.want) {
			t.Errorf("unwated result! got:%v; want:%v", gotSlice, test.want)
		}
	}
}
