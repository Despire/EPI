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

import "testing"

func TestToBase10(t *testing.T) {
	tests := []struct {
		input string
		base  int
		want  int
	}{
		{
			input: "01",
			base:  2,
			want:  1,
		},

		{
			input: "11",
			base:  2,
			want:  3,
		},

		{
			input: "1A",
			base:  16,
			want:  26,
		},

		{
			input: "AC",
			base:  16,
			want:  172,
		},

		{
			input: "102",
			base:  10,
			want:  102,
		},
	}

	for _, test := range tests {
		got := toBase10(test.input, test.base)
		if got != test.want {
			t.Errorf("Unwanted result! got value: %v, wanted value: %v", got, test.want)
		}
	}
}

func TestConvertFromBase10(t *testing.T) {
	tests := []struct {
		input int
		base  int
		want  string
	}{
		{
			input: 10,
			base:  16,
			want:  "A",
		},

		{
			input: 18,
			base:  16,
			want:  "12",
		},

		{
			input: 4,
			base:  2,
			want:  "100",
		},
	}

	for _, test := range tests {
		got := convertFromBase10(test.input, test.base)

		if got != test.want {
			t.Errorf("Unwanted result! got values: %v, wanted value: %v", got, test.want)
		}
	}
}

func TestConvertBase(t *testing.T) {
	tests := []struct {
		input string
		b1    int
		b2    int
		want  string
	}{
		{
			input: "-127",
			b1:    10,
			b2:    2,
			want:  "-1111111",
		},

		{
			input: "-1111111",
			b1:    2,
			b2:    16,
			want:  "-7F",
		},
	}

	for _, test := range tests {
		got := convertBase(test.input, test.b1, test.b2)

		if got != test.want {
			t.Errorf("Unwanted result! got values: %v, wanted value: %v", got, test.want)
		}
	}
}
