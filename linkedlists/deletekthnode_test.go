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

func TestDeleteKthNode(t *testing.T) {
	tests := []struct {
		input    *Node
		k        int
		want     []int
		wantBool bool
	}{
		{
			input:    nil,
			k:        5,
			want:     make([]int, 0),
			wantBool: false,
		},

		{
			input: &Node{
				Data: 5,
				Next: nil,
			},
			k:        2,
			wantBool: false,
			want:     []int{5},
		},
		{
			input: &Node{
				Data: 5,
				Next: &Node{
					Data: 10,
					Next: nil,
				},
			},
			k:        3,
			wantBool: false,
			want:     []int{5, 10},
		},
		{
			input: &Node{
				Data: 5,
				Next: &Node{
					Data: 10,
					Next: &Node{
						Data: 15,
						Next: nil,
					},
				},
			},
			k:        1,
			wantBool: true,
			want:     []int{5, 10},
		},
		{
			input: &Node{
				Data: 5,
				Next: &Node{
					Data: 10,
					Next: &Node{
						Data: 15,
						Next: &Node{
							Data: 20,
							Next: &Node{
								Data: 25,
								Next: nil,
							},
						},
					},
				},
			},
			k:        3,
			wantBool: true,
			want:     []int{5, 10, 20, 25},
		},
	}

	for _, test := range tests {
		got := DeleteKthLast(test.input, test.k)

		if got != test.wantBool {
			t.Errorf("unwated result! got: %v, want: %v", got, test.wantBool)
		}

		order := make([]int, 0)
		for test.input != nil {
			order = append(order, test.input.Data.(int))
			test.input = test.input.Next
		}

		if !reflect.DeepEqual(order, test.want) {
			t.Errorf("unwanted result! got: %v, want: %v", order, test.want)
		}
	}

}
