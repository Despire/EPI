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

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		input *Node
		want  []int
	}{
		{
			input: &Node{
				Data: 5,
				Next: &Node{
					Data: 10,
					Next: &Node{
						Data: 10,
						Next: &Node{
							Data: 10,
							Next: nil,
						},
					},
				},
			},

			want: []int{5, 10},
		},

		{
			input: &Node{
				Data: 5,
				Next: nil,
			},
			want: []int{5},
		},

		{
			input: &Node{
				Data: 2,
				Next: &Node{
					Data: 2,
					Next: &Node{
						Data: 3,
						Next: &Node{
							Data: 4,
							Next: &Node{
								Data: 4,
								Next: nil,
							},
						},
					},
				},
			},
			want: []int{2, 3, 4},
		},
	}

	for _, test := range tests {
		RemoveDuplicates(test.input)

		got := []int{}

		for test.input != nil {
			got = append(got, test.input.Data.(int))
			test.input = test.input.Next
		}

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("unwated result! got: %v, want: %v", got, test.want)
		}
	}
}
