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

func TestHasCycle(t *testing.T) {
	tests := []struct {
		input *Node
		want  bool
	}{
		{
			input: &Node{
				Data: 5,
				Next: &Node{
					Data: 10,
					Next: &Node{
						Data: 20,
						Next: &Node{
							Data: 50,
							Next: &Node{
								Data: 60,
								Next: &Node{
									Data: 70,
									Next: nil,
								},
							},
						},
					},
				},
			},

			want: true,
		},
		{
			input: &Node{
				Data: 5,
				Next: &Node{
					Data: 10,
					Next: &Node{
						Data: 20,
						Next: &Node{
							Data: 50,
							Next: &Node{
								Data: 60,
								Next: &Node{
									Data: 70,
									Next: nil,
								},
							},
						},
					},
				},
			},

			want: false,
		},
		{
			input: &Node{
				Data: 10,
				Next: nil,
			},
			want: false,
		},
		{
			input: &Node{
				Data: 10,
				Next: &Node{
					Data: 20,
					Next: nil,
				},
			},
			want: false,
		},
	}

	tests[0].input.Next.Next.Next.Next.Next.Next = tests[0].input.Next.Next.Next

	for _, test := range tests {
		got := HasCycle(test.input)

		if got != test.want {
			t.Errorf("unwated result! got: %v; want: %v", got, test.want)
		}
	}
}

func TestRootOfClycle(t *testing.T) {
	tests := []struct {
		input *Node
		want  *Node
	}{
		{
			input: &Node{
				Data: 5,
				Next: &Node{
					Data: 10,
					Next: &Node{
						Data: 20,
						Next: &Node{
							Data: 50,
							Next: &Node{
								Data: 60,
								Next: &Node{
									Data: 70,
									Next: nil,
								},
							},
						},
					},
				},
			},
		},
		{
			input: &Node{
				Data: 5,
				Next: &Node{
					Data: 10,
					Next: &Node{
						Data: 20,
						Next: &Node{
							Data: 50,
							Next: &Node{
								Data: 60,
								Next: &Node{
									Data: 70,
									Next: nil,
								},
							},
						},
					},
				},
			},
		},
		{
			input: &Node{
				Data: 10,
				Next: nil,
			},
		},
		{
			input: &Node{
				Data: 10,
				Next: &Node{
					Data: 20,
					Next: nil,
				},
			},
		},
	}

	tests[0].input.Next.Next.Next.Next.Next.Next = tests[0].input.Next.Next.Next
	tests[0].want = tests[0].input.Next.Next.Next

	for _, test := range tests {
		got := RootOfCycle(test.input)

		if got != test.want {
			t.Errorf("unwated result! got: %v; want: %v", got, test.want)
		}
	}
}
