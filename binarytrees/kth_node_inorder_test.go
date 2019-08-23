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

package binarytrees

import "testing"

func TestKthNodeInorderTraversal(t *testing.T) {
	root := &Node{
		Data: 4,
		Left: &Node{
			Data: 1,
			Right: &Node{
				Data: 3,
				Left: &Node{
					Data: 2,
				},
			},
		},
		Right: &Node{
			Data: 6,
			Left: &Node{
				Data: 5,
			},
			Right: &Node{
				Data: 7,
			},
		},
	}

	tests := []struct {
		root *Node

		k    int
		want *Node
	}{
		{
			root: root,
			k:    6,
			want: root.Right,
		},
		{
			root: root,
			k:    3,
			want: root.Left.Right,
		},
		{
			root: root,
			k:    1,
			want: root.Left,
		},
		{
			root: root,
			k:    5,
			want: root.Right.Left,
		},
		{
			root: nil,
			k:    5,
			want: nil,
		},
		{
			root: root,
			k:    10,
			want: nil,
		},
	}

	for _, test := range tests {
		have := KthNodeInorderTraversal(test.k, test.root)
		if have != test.want {
			t.Errorf("result mismatch, have: %v, want: %v", have, test.want)
		}
	}
}

func TestKthNodeInorderTraversalWithModifiedNode(t *testing.T) {
	root := &ModifiedNode{
		Data:        4,
		subtreeSize: 7,
		Left: &ModifiedNode{
			Data:        1,
			subtreeSize: 3,
			Right: &ModifiedNode{
				Data:        3,
				subtreeSize: 2,
				Left: &ModifiedNode{
					Data:        2,
					subtreeSize: 1,
				},
			},
		},
		Right: &ModifiedNode{
			Data:        6,
			subtreeSize: 3,
			Left: &ModifiedNode{
				Data:        5,
				subtreeSize: 1,
			},
			Right: &ModifiedNode{
				Data:        7,
				subtreeSize: 1,
			},
		},
	}

	tests := []struct {
		root *ModifiedNode

		k    int
		want *ModifiedNode
	}{
		{
			root: root,
			k:    6,
			want: root.Right,
		},
		{
			root: root,
			k:    3,
			want: root.Left.Right,
		},
		{
			root: root,
			k:    1,
			want: root.Left,
		},
		{
			root: root,
			k:    5,
			want: root.Right.Left,
		},
		{
			root: nil,
			k:    5,
			want: nil,
		},
		{
			root: root,
			k:    10,
			want: nil,
		},
		{
			root: root,
			k:    0,
			want: nil,
		},
	}

	for _, test := range tests {
		have := KthNodeInorderTraversalWithModifiedNode(test.k, test.root)
		if have != test.want {
			t.Errorf("result mismatch, have: %v, want: %v", have, test.want)
		}
	}

}
