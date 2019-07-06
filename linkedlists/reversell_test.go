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

func TestReverse(t *testing.T) {
	tests := []struct {
		input *Node
		want  []int
	}{
		{
			input: nil,
			want:  make([]int, 0),
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
				Data: 5,
				Next: &Node{
					Data: 10,
					Next: nil,
				},
			},
			want: []int{10, 5},
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
			want: []int{15, 10, 5},
		},
	}

	for _, test := range tests {
		got := Reverse(test.input)

		order := make([]int, 0)

		for got != nil {
			order = append(order, got.Data.(int))
			got = got.Next
		}

		if !reflect.DeepEqual(order, test.want) {
			t.Errorf("unwanted result! got: %v, want: %v", order, test.want)
		}
	}
}

func TestReverseSubList(t *testing.T) {
	tests := []struct {
		input *Node
		i, j  int
		want  []int
	}{
		{
			input: nil,
			want:  make([]int, 0),
		},

		{
			input: &Node{
				Data: 5,
				Next: nil,
			},
			want: []int{5},
		},
		{
			i: 0,
			j: 1,
			input: &Node{
				Data: 5,
				Next: &Node{
					Data: 10,
					Next: nil,
				},
			},
			want: []int{10, 5},
		},
		{
			i: 2,
			j: 3,
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
			want: []int{5, 15, 10},
		},
		{
			i: 2,
			j: 4,
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
			want: []int{5, 20, 15, 10, 25},
		},
	}

	for _, test := range tests {
		got := ReverseSubList(test.input, test.i, test.j)

		order := make([]int, 0)

		for got != nil {
			order = append(order, got.Data.(int))
			got = got.Next
		}

		if !reflect.DeepEqual(order, test.want) {
			t.Errorf("unwanted result! got: %v, want: %v", order, test.want)
		}
	}
}
