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

func TestReconstructFromPreorder(t *testing.T) {
	tests := []struct {
		preorder []string

		wantPreOrder []string
		wantInOrder  []string
	}{
		{
			preorder:     []string{"H", "B", "null", "E", "null", "null", "C", "null", "D", "null", "null"},
			wantPreOrder: []string{"H", "B", "E", "C", "D"},
			wantInOrder:  []string{"B", "E", "H", "C", "D"},
		},
		{
			preorder:     []string{"H", "B", "F", "null", "null", "E", "A", "null", "null", "null", "C", "null", "D", "null", "G", "I", "null", "null", "null"},
			wantPreOrder: []string{"H", "B", "F", "E", "A", "C", "D", "G", "I"},
			wantInOrder:  []string{"F", "B", "A", "E", "H", "C", "D", "I", "G"},
		},
		{
			preorder:     nil,
			wantInOrder:  nil,
			wantPreOrder: nil,
		},
		{
			preorder:     []string{"1", "2", "null", "null", "3", "null", "null"},
			wantInOrder:  []string{"2", "1", "3"},
			wantPreOrder: []string{"1", "2", "3"},
		},
		{
			preorder:     []string{"1", "2", "4", "null", "null", "5", "null", "null", "3", "6", "null", "null", "null"},
			wantPreOrder: []string{"1", "2", "4", "5", "3", "6"},
			wantInOrder:  []string{"4", "2", "5", "1", "6", "3",},
		},
		{
			preorder:     []string{"1", "2", "4", "null", "null", "null", "3", "null", "null"},
			wantPreOrder: []string{"1", "2", "4", "3"},
			wantInOrder:  []string{"4", "2", "1", "3"},
		},
		{
			preorder:     []string{"1", "2", "3", "null", "null", "null", "null"},
			wantPreOrder: []string{"1", "2", "3"},
			wantInOrder:  []string{"3", "2", "1"},
		},
	}

	for _, test := range tests {
		have := ReconstructFromPreorder(test.preorder)

		havePreorder := IterativePreorder(have)
		haveInorder := IterativeInorder(have)

		for i := range test.wantPreOrder {
			if test.wantPreOrder[i] != havePreorder[i].Data.(string) {
				t.Errorf("result mismatch! have:%v; want:%v", havePreorder[i].Data.(string), test.wantPreOrder[i])
			}
		}
		for i := range test.wantInOrder {
			if test.wantInOrder[i] != haveInorder[i].Data.(string) {
				t.Errorf("result mismatch! have:%v; want:%v", haveInorder[i].Data.(string), test.wantInOrder[i])
			}
		}
	}
}
