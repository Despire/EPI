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

func TestReconstruct(t *testing.T) {
	tests := []struct {
		inorder  []string
		preorder []string
	}{
		{
			inorder:  []string{"F", "B", "A", "E", "H", "C", "D", "I", "G"},
			preorder: []string{"H", "B", "F", "E", "A", "C", "D", "G", "I"},
		},

		{
			inorder:  nil,
			preorder: nil,
		},
		{
			inorder: []string{"Q", "B", "E"},
			preorder: []string{"B", "Q", "E"},
		},
	}

	for _, test := range tests {
		have := Reconstruct(test.preorder, test.inorder)

		inorderHave := IterativeInorder(have)
		preorderHave := IterativePreorder(have)

		for i := range inorderHave {
			if inorderHave[i].Data.(string) != test.inorder[i] {
				t.Errorf("result mismatch! have: %v, want: %v", inorderHave[i].Data.(string), test.inorder[i])
			}
		}
		for i := range preorderHave {
			if preorderHave[i].Data.(string) != test.preorder[i] {
				t.Errorf("result mismatch! have: %v, want: %v", preorderHave[i].Data.(string), test.preorder[i])
			}
		}
	}
}
