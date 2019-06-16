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

func TestReverseWords(t *testing.T) {
	tests := []struct {
		input []rune
		want  []rune
	}{
		{
			input: []rune("Alice likes Bob"),
			want:  []rune("Bob likes Alice"),
		},

		{
			input: []rune("Aloha"),
			want:  []rune("Aloha"),
		},

		{
			input: []rune("Aloha bamamas"),
			want:  []rune("bamamas Aloha"),
		},

		{
			input: []rune("古籍書寫中 同一個字的不同寫法"),
			want:  []rune("同一個字的不同寫法 古籍書寫中"),
		},
	}

	for _, test := range tests {
		ReverseWords(test.input)
		if !reflect.DeepEqual(test.input, test.want) {
			t.Errorf("Unexpected result! got: %v; want: %v", test.input, test.want)
		}
	}
}
