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

package linkedlists

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input *Node
		want  bool
	}{
		{
			input: &Node{
				Data: 1,
				Next: &Node{
					Data: 2,
					Next: &Node{
						Data: 3,
						Next: &Node{
							Data: 2,
							Next: &Node{
								Data: 1,
								Next: nil,
							},
						},
					},
				},
			},
			want: true,
		},
		{
			input: &Node{
				Data: 1,
				Next: &Node{
					Data: 2,
					Next: &Node{
						Data: 3,
						Next: nil,
					},
				},
			},
			want: false,
		},

		{
			input: nil,
			want:  false,
		},

		{
			input: &Node{
				Data: 1,
				Next: &Node{
					Data: 2,
					Next: nil,
				},
			},
			want: false,
		},
	}

	for i, test := range tests {
		got := IsPalindrome(test.input)

		if got != test.want {
			t.Errorf("sequence: %d", i)
			t.Errorf("unwanted result! got:%v; want:%v", got, test.want)
		}
	}
}
