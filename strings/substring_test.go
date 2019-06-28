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
	"testing"
)

func TestHash(t *testing.T) {
	tests := []struct {
		input     string
		base      uint64
		want      uint64
		wantPower uint64
	}{
		{
			input:     "hello",
			base:      26,
			want:      49376607,
			wantPower: 4,
		},

		{
			input:     "world",
			base:      26,
			want:      56411052,
			wantPower: 4,
		},
	}

	for _, test := range tests {
		gotResult, gotPower := hash(test.input, test.base)

		if gotResult != test.want {
			t.Errorf("unwanted result! got: %v, want: %v", gotResult, test.want)
		}
		if gotPower != test.wantPower {
			t.Errorf("unwanted result! got: %v, want: %v", gotPower, test.wantPower)
		}

	}
}

func TestFindFirstOf(t *testing.T) {
	tests := []struct {
		input string
		sub   string
		want  int
	}{
		{
			input: "hello world",
			sub:   "llo",
			want:  2,
		},

		{
			input: "hello world",
			sub:   "rld",
			want:  8,
		},

		{
			input: "hello world",
			sub:   " ",
			want:  5,
		},

		{
			input: "hello world",
			sub:   "e",
			want:  1,
		},
		{
			input: "hello world",
			sub:   "d",
			want:  10,
		},
	}

	for _, test := range tests {
		got := findFirstOf(test.sub, test.input)

		if got != test.want {
			t.Errorf("unwanted result! got: %v, want: %v", got, test.want)
		}
	}
}
