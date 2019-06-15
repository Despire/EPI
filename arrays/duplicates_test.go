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

package arrays

import (
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	tests := []struct {
		in   []int
		want int
	}{
		{[]int{2, 2, 3, 4, 4, 5}, 4},
		{[]int{2, 3, 4, 4, 7, 11, 11, 11, 13}, 6},
		{[]int{}, 0},
		{[]int{1}, 1},
		{nil, 0},
		{[]int{1}, 1},
		{[]int{1, 1}, 1},
		{[]int{1, 2, 3}, 3},
		{[]int{1, 1, 2, 2, 3, 3}, 3},
		{[]int{2, 3, 5, 5, 7, 11, 11, 11, 13}, 6},
	}
	for _, test := range tests {
		got := DeleteDuplicates(test.in)
		if got != test.want {
			t.Errorf("want %v; got %v", test.want, got)
		}
	}
}
