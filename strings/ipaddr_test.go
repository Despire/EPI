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
// limitations under the Licens

package strings

import (
	"reflect"
	"testing"
)

func TestComputIPAddr(t *testing.T) {
	tests := []struct {
		input string
		want  []string
	}{
		{
			input: "19216811",
			want: []string{
				"1.92.168.11",
				"19.2.168.11",
				"19.21.68.11",
				"19.216.8.11",
				"19.216.81.1",
				"192.1.68.11",
				"192.16.8.11",
				"192.16.81.1",
				"192.168.1.1",
			},
		},

		{
			input: "1111",
			want: []string{
				"1.1.1.1",
			},
		},

		{
			input: "19111",
			want: []string{
				"1.9.1.11",
				"1.9.11.1",
				"1.91.1.1",
				"19.1.1.1",
			},
		},
	}

	for _, test := range tests {
		got := ComputeIpv4(test.input)

		if !reflect.DeepEqual(test.want, got) {
			t.Errorf("unwated result! got; %v; want: %v", got, test.want)
		}
	}
}
