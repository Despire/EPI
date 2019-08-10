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

import "strings"

// IsWellFormed returns true if the
// brackets in the strings are well formed.
// Time complexity O(N).
func IsWellFormed(s string) bool {
	leftBrackets := "{[("
	rightBrackets := "}])"
	stack := new(stack)

	for _, r := range s {
		if strings.ContainsRune(leftBrackets, r) {
			stack.Push(r)
		} else if strings.ContainsRune(rightBrackets, r) {
			if stack.Size() == 0 {
				return false
			}
			leftBracket := stack.Pop().(rune)
			switch r {
			case '}':
				if leftBracket != '{' {
					return false
				}
			case ']':
				if leftBracket != '[' {
					return false
				}
			case ')':
				if leftBracket != '(' {
					return false
				}
			}
		}
	}

	return stack.Size() == 0
}
