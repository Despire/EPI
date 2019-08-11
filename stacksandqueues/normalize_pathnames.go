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
	"errors"
	"strings"
)

var (
	// ErrInvalidPath is returned if the path argument
	// to the NormalizePath function is invalid.
	ErrInvalidPath = errors.New("path is invalid")
)

// NormalizePath takes a string that is a path
// to a file/folder and returns a string that
// is the shortest equivalent pathname.
// Time complexity O(N).
func NormalizePath(path string) (string, error) {
	if path == "" {
		return "", ErrInvalidPath
	}
	const (
		currDir   = "."
		parentDir = ".."
		split     = "/"
	)

	stack := make([]string, 0)
	// absolute path is a pecial case where '/' is considered a directory
	if strings.HasPrefix(path, split) {
		stack = append(stack, split)
		path = path[len(split):]
	}

	for _, token := range strings.Split(path, split) {
		switch token {
		case currDir:
			// do nothing we skip the current dir.
		case parentDir:
			// since we don't know the current directory, let's ignore
			// the case where we can get from ../ to the current dir.
			// we handle only the case where we can get up from this dir

			// cases where the path starts with ../../
			// or if we get up from the curr dit e.g ./example/../../../ -> ../../
			if len(stack) == 0 || stack[len(stack)-1] == parentDir {
				stack = append(stack, token)
			} else {
				// if we try to go up one dir from root.
				if stack[len(stack)-1] == split {
					return "", ErrInvalidPath
				}
				stack = stack[:len(stack)-1]
			}
		default:
			if token != "" {
				stack = append(stack, token)
			}
		}
	}

	stringer := new(strings.Builder)

	if len(stack) != 0 && stack[0] == split {
		stringer.WriteString(split)
		stack = stack[1:]
	}

	for i, token := range stack {
		stringer.WriteString(token)

		if i != len(stack)-1 {
			stringer.WriteString(split)
		}
	}

	return stringer.String(), nil
}
