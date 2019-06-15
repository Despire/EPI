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

package strings

import (
	"reflect"
	"testing"
)

func TestReplaceAndRemove(t *testing.T) {
	tests := []struct {
		input []rune
		size  int
		want  []rune
	}{
		{
			input: []rune{'a', 'c', 'd', 'b', 'b', 'c', 'a'},
			size:  7,
			want:  []rune{'d', 'd', 'c', 'd', 'c', 'd', 'd'},
		},

		{
			input: []rune{'a', 'b', 'a', 'c', ' '},
			size:  4,
			want:  []rune{'d', 'd', 'd', 'd', 'c'},
		},

		{
			input: []rune{'b', 'a', 'a', 'c', ' '},
			size:  4,
			want:  []rune{'d', 'd', 'd', 'd', 'c'},
		},

		{
			input: []rune{'b', 'b', 'a', 'c', 'b', 'a', 'a'},
			size:  7,
			want:  []rune{'d', 'd', 'c', 'd', 'd', 'd', 'd'},
		},
	}

	for _, test := range tests {
		ReplaceAndRemove(test.input, test.size)

		if !reflect.DeepEqual(test.input, test.want) {
			t.Errorf("Unexpected result! got: %v; want:%v", test.input, test.want)
		}
	}
}
