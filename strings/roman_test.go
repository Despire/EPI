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

package strings

import "testing"

func TestRomanToDecimal(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: "IV",
			want:  4,
		},

		{
			input: "IX",
			want:  9,
		},
		{
			input: "XL",
			want:  40,
		},
		{
			input: "LIX",
			want:  59,
		},
	}

	for _, test := range tests {
		got := RomanToDecimal(test.input)

		if got != test.want {
			t.Errorf("unwanted result ! got: %v; want: %v", got, test.want)
		}

	}
}
