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

func TestDecodeEncoding(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: "A",
			want:  1,
		},

		{
			input: "AA",
			want:  27,
		},

		{
			input: "ZZ",
			want:  702,
		},
	}

	for _, test := range tests {
		got := decodeEncoding(test.input)

		if got != test.want {
			t.Errorf("Unexpected value! got: %v; want: %v", got, test.want)
		}
	}
}