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

func TestPartition(t *testing.T) {
	tests := []struct {
		i    int
		in   [5]int
		want [5]int
	}{
		{0, [...]int{5, 4, 3, 2, 1}, [...]int{4, 3, 2, 1, 5}},
		{1, [...]int{5, 4, 3, 2, 1}, [...]int{1, 3, 2, 4, 5}},
		{2, [...]int{5, 4, 3, 2, 1}, [...]int{1, 2, 3, 4, 5}},
		{3, [...]int{5, 4, 3, 2, 1}, [...]int{1, 2, 3, 4, 5}},
		{1, [...]int{3, 3, 3, 2, 1}, [...]int{2, 1, 3, 3, 3}},
	}

	for _, test := range tests {
		got := test.in
		e1, e2 := Partition(got[:], test.i)

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("want %v; got %v", test.want, got)
		}

		pivot := test.in[test.i]
		for i := 0; i < e1; i++ {
			if !(got[i] < pivot) {
				t.Errorf("elements that are less than the pivot are not correct")
			}
		}
		for i := e1; i < e2; i++ {
			if !(got[i] == pivot) {
				t.Errorf("elements that are equal to the pivot are not correct")
			}
		}
		for i := e2; i < len(got); i++ {
			if !(got[i] > pivot) {
				t.Errorf("elements that are greater than the pivot are not correct")
			}
		}
	}
}
