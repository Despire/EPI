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

func TestMerge(t *testing.T) {
	tests := []struct {
		inputFirst  *Node
		inputSecond *Node
		want        []int
	}{
		{
			inputFirst: &Node{
				Data: 3,
				Next: &Node{
					Data: 4,
					Next: &Node{
						Data: 5,
						Next: nil,
					},
				},
			},

			inputSecond: &Node{
				Data: 1,
				Next: &Node{
					Data: 3,
					Next: &Node{
						Data: 6,
						Next: nil,
					},
				},
			},

			want: []int{1, 3, 3, 4, 5, 6},
		},

		{
			inputFirst:  nil,
			inputSecond: nil,
			want:        make([]int, 0),
		},
	}

	for _, test := range tests {
		got := Merge(test.inputFirst, test.inputSecond)

		gotSlice := make([]int, 0)

		for got != nil {
			gotSlice = append(gotSlice, got.Data.(int))
			got = got.Next
		}

		if !reflect.DeepEqual(gotSlice, test.want) {
			t.Errorf("unwanted result! got: %v, wat: %v", gotSlice, test.want)
		}
	}
}
