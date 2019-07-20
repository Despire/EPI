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

func TestShiftRightByK(t *testing.T) {
	tests := []struct {
		input *Node
		k     int
		want  []int
	}{
		{
			input: &Node{
				Data: 2,
				Next: &Node{
					Data: 3,
					Next: &Node{
						Data: 5,
						Next: &Node{
							Data: 3,
							Next: &Node{
								Data: 2,
								Next: nil,
							},
						},
					},
				},
			},
			k:    3,
			want: []int{5, 3, 2, 2, 3},
		},
		{
			input: nil,
			k:     3,
			want:  []int{},
		},
		{
			input: &Node{
				Data: 2,
				Next: &Node{
					Data: 3,
					Next: &Node{
						Data: 5,
						Next: &Node{
							Data: 3,
							Next: &Node{
								Data: 2,
								Next: nil,
							},
						},
					},
				},
			},
			k:    4,
			want: []int{3, 5, 3, 2, 2},
		},
		{
			input: &Node{
				Data: 2,
				Next: &Node{
					Data: 3,
					Next: &Node{
						Data: 5,
						Next: &Node{
							Data: 3,
							Next: &Node{
								Data: 2,
								Next: nil,
							},
						},
					},
				},
			},
			k:    5,
			want: []int{2, 3, 5, 3, 2},
		},
		{
			input: &Node{
				Data: 2,
				Next: &Node{
					Data: 3,
					Next: &Node{
						Data: 5,
						Next: &Node{
							Data: 3,
							Next: &Node{
								Data: 2,
								Next: nil,
							},
						},
					},
				},
			},
			k:    1,
			want: []int{2, 2, 3, 5, 3},
		},
	}

	for _, test := range tests {
		gotRoot := ShiftRightByK(test.input, test.k)

		gotSlice := []int{}
		for gotRoot != nil {
			gotSlice = append(gotSlice, gotRoot.Data.(int))
			gotRoot = gotRoot.Next
		}

		if !reflect.DeepEqual(gotSlice, test.want) {
			t.Errorf("unwatned result! got:%v; want:%v", gotSlice, test.want)
		}
	}
}
