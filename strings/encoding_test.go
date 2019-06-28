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

func TestRLE(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "aaaabcccaa",
			want:  "4a1b3c2a",
		},
		{
			input: "eeeffffee",
			want:  "3e4f2e",
		},
		{
			input: "aaaaaaaaaaabbb",
			want:  "11a3b",
		},
	}

	for _, test := range tests {
		got := RLE(test.input)

		if got != test.want {
			t.Errorf("unwated result; got: %v; want: %v", got, test.want)
		}
	}
}

func TestRLD(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "4a1b3c2a",
			want:  "aaaabcccaa",
		},
		{
			input: "3e4f2e",
			want:  "eeeffffee",
		},
		{
			input: "11a3b",
			want:  "aaaaaaaaaaabbb",
		},
	}

	for _, test := range tests {
		got := RLD(test.input)

		if got != test.want {
			t.Errorf("unwated result; got: %v; want: %v", got, test.want)
		}
	}
}
