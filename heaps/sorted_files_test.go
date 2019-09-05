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
// limitations under the License.

package heaps

import (
	"reflect"
	"testing"
)

func TestMergeFiles(t *testing.T) {
	tests := []struct {
		in   [][]int
		want [] int
		err  error
	}{
		{
			in: [][]int{
				{3, 5, 6},
				{2, 4, 5},
				{5, 7, 9},
			},

			want: []int{
				2, 3, 4, 5, 5, 5, 6, 7, 9,
			},

			err: nil,
		},

		{
			in: [][]int{
				{1, 4, 7},
				{2, 3},
				{5},
				{7, 9, 10},
				{3, 5, 6},
			},
			want: []int{
				1, 2, 3, 3, 4, 5, 5, 6, 7, 7, 9, 10,
			},
			err: nil,
		},
		{
			in:   nil,
			want: []int{},
			err:  nil,
		},
	}

	for _, test := range tests {
		have, err := MergeFiles(test.in...)

		if err != test.err {
			t.Errorf("error result mismatch! have: %v; want: %v", err, test.err)
		}

		if !reflect.DeepEqual(have, test.want) {
			t.Errorf("slice result mismach! have: %v; want: %v", have, test.want)
		}
	}

}
