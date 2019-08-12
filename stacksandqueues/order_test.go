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

package stacksandqueues

import (
	"reflect"
	"testing"

	"github.com/despire/EPI/binarytrees"
)

func TestOrderOfIncreasingDepth(t *testing.T) {
	tests := []struct {
		// in
		root *binarytrees.Node
		// out
		want [][]interface{}
	}{
		{
			root: &binarytrees.Node{
				Data: 314,
				Left: &binarytrees.Node{
					Data: 6,
					Left: &binarytrees.Node{
						Data: 271,
						Left: &binarytrees.Node{
							Data:  28,
							Left:  nil,
							Right: nil,
						},
						Right: &binarytrees.Node{
							Data:  0,
							Left:  nil,
							Right: nil,
						},
					},
					Right: &binarytrees.Node{
						Data: 561,
						Left: nil,
						Right: &binarytrees.Node{
							Data: 3,
							Left: &binarytrees.Node{
								Data:  17,
								Left:  nil,
								Right: nil,
							},
							Right: nil,
						},
					},
				},
				Right: &binarytrees.Node{
					Data: 6,
					Left: &binarytrees.Node{
						Data: 2,
						Left: nil,
						Right: &binarytrees.Node{
							Data: 1,
							Left: &binarytrees.Node{
								Data: 401,
								Left: nil,
								Right: &binarytrees.Node{
									Data:  641,
									Left:  nil,
									Right: nil,
								},
							},
							Right: &binarytrees.Node{
								Data:  257,
								Left:  nil,
								Right: nil,
							},
						},
					},
					Right: &binarytrees.Node{
						Data: 271,
						Left: nil,
						Right: &binarytrees.Node{
							Data:  28,
							Left:  nil,
							Right: nil,
						},
					},
				},
			},
			want: [][]interface{}{
				{314},
				{6, 6},
				{271, 561, 2, 271},
				{28, 0, 3, 1, 28},
				{17, 401, 257},
				{641},
			},
		},
		{
			root: nil,
			want: nil,
		},
		{
			root: &binarytrees.Node{
				Data: 314,
				Left: &binarytrees.Node{
					Data:  5,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
			want: [][]interface{}{
				{314},
				{5},
			},
		},
	}

	for _, test := range tests {
		got := OrderOfIncreasingDepth(test.root)

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("result mismatch. got: %v; want: %v", got, test.want)
		}
	}
}
