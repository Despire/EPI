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
	"reflect"
	"testing"
)

func TestApplePermutation(t *testing.T) {
	tests := []struct {
		in   []int
		perm []int
		want []int
	}{
		{[]int{0, 1, 2, 3}, []int{2, 0, 1, 3}, []int{1, 2, 0, 3}},
		{[]int{0, 1, 2, 3}, []int{0, 2, 1, 3}, []int{0, 2, 1, 3}},
		{[]int{0, 1, 2, 3}, []int{0, 1, 2, 3}, []int{0, 1, 2, 3}},
		{[]int{0, 1, 2, 3}, []int{3, 2, 1, 0}, []int{3, 2, 1, 0}},
		{[]int{0, 1, 2, 3}, []int{3, 1, 2, 0}, []int{3, 1, 2, 0}},
	}

	for _, test := range tests {
		ApplyPermutation(test.in, test.perm)
		if !reflect.DeepEqual(test.want, test.in) {
			t.Errorf("want %v, got %v", test.want, test.in)
		}
	}
}
