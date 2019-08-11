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

import "math"

// hash hashes the string s, using base 10.
// Takes O(N) time with O(1) space complexity.
func hash(s string, b uint64) (uint64, uint64) {
	result := uint64(0)
	power := int64(-1)

	for i := 0; i < len(s); i++ {
		result = result*b + uint64(s[i])
		power++
	}

	return result, uint64(power)
}

// findFirstOf finds the first occurence of the substring <sub>
// in the string <in>.
// Takes O(m + n) time with O(1) space complexity.
func findFirstOf(sub string, in string) int {
	const (
		base = 26
	)

	subHash, _ := hash(sub, base)

	inSub := in[:len(sub)]

	inHash, inPower := hash(inSub, base)

	for i := len(sub); i < len(in); i++ {
		if subHash == inHash && sub == inSub {
			return i - len(sub)
		}

		inHash -= uint64(math.Pow(float64(base), float64(inPower))) * uint64(inSub[0])
		inHash = inHash*base + uint64(in[i])
		inSub = in[(i+1)-len(sub) : i+1]
	}
	if subHash == inHash && sub == inSub {
		return len(in) - len(sub)
	}

	return 0
}
